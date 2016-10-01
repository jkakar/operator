package bread

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/sr/operator"
	"github.com/sr/operator/hipchat"
	"golang.org/x/net/context"
	"golang.org/x/net/context/ctxhttp"
)

type canoeDeployer struct {
	operator.Replier
	canoeURL string
	apiToken string
	http     *http.Client
}

type canoeProject struct {
	Name string `json:"name"`
}

type canoeBuild struct {
	ArtifactURL string    `json:"artifact_url"`
	RepoURL     string    `json:"repo_url"`
	URL         string    `json:"url"`
	Branch      string    `json:"branch"`
	BuildNumber int       `json:"build_number"`
	SHA         string    `json:"sha"`
	PassedCI    bool      `json:"passed_ci"`
	CreatedAt   time.Time `json:"created_at"`
}

func (d *canoeDeployer) listTargets(ctx context.Context) (targets []*DeployTarget, err error) {
	var resp *http.Response
	if resp, err = d.doCanoe(ctx, "GET", "/api/projects", ""); err == nil {
		defer func() { _ = resp.Body.Close() }()
		var projs []*canoeProject
		if err := json.NewDecoder(resp.Body).Decode(&projs); err == nil {
			for _, p := range projs {
				targets = append(targets, &DeployTarget{Name: p.Name, Canoe: true})
			}
		}
	}
	return targets, err
}

func (d *canoeDeployer) listBuilds(ctx context.Context, t *DeployTarget, branch string) ([]build, error) {
	resp, err := d.doCanoe(
		ctx,
		"GET",
		fmt.Sprintf("/api/projects/%s/branches/%s/builds", t.Name, branch),
		"",
	)
	defer func() { _ = resp.Body.Close() }()
	if err != nil {
		return nil, err
	}
	var data []*canoeBuild
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	var builds []build
	for _, b := range data {
		builds = append(builds, build(b))
	}
	return builds, nil
}

func (d *canoeDeployer) deploy(ctx context.Context, req *operator.Request, t *DeployTarget, reqBuild build) (*operator.Message, error) {
	params := url.Values{}
	params.Add("project_name", t.Name)
	params.Add("artifact_url", reqBuild.GetURL())
	params.Add("user_email", req.UserEmail())
	resp, err := d.doCanoe(
		ctx,
		"POST",
		"/api/targets/production/deploys",
		params.Encode(),
	)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()
	type canoeDeploy struct {
		ID int `json:"id"`
	}
	type canoeResp struct {
		Error   bool         `json:"error"`
		Message string       `json:"message"`
		Deploy  *canoeDeploy `json:"deploy"`
	}
	var data canoeResp
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	if data.Error {
		return nil, errors.New(data.Message)
	}
	deployURL := fmt.Sprintf("%s/projects/%s/deploys/%d?watching=1", d.canoeURL, t.Name, data.Deploy.ID)
	return &operator.Message{
		Text: deployURL,
		HTML: fmt.Sprintf(`Deployment of %s triggered. Watch it here: <a href="%s">#%d</a>`, t.Name, deployURL, data.Deploy.ID),
		Options: &operatorhipchat.MessageOptions{
			Color: "green",
		},
	}, nil
}

func (d *canoeDeployer) doCanoe(ctx context.Context, meth, path, body string) (*http.Response, error) {
	req, err := http.NewRequest(meth, d.canoeURL+path, strings.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-Api-Token", d.apiToken)
	if body != "" {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}
	resp, err := ctxhttp.Do(ctx, d.http, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		if body, err := ioutil.ReadAll(resp.Body); err == nil {
			return nil, fmt.Errorf("canoe API request failed with status %d and body: %s", resp.StatusCode, body)
		}
		return nil, fmt.Errorf("canoe API request failed with status %d", resp.StatusCode)
	}
	return resp, nil
}

func (a *canoeBuild) GetNumber() int {
	if a == nil {
		return 0
	}
	return a.BuildNumber
}

func (a *canoeBuild) GetBambooID() string {
	if a == nil {
		return ""
	}
	return ""
}

func (a *canoeBuild) GetBranch() string {
	if a == nil {
		return ""
	}
	return a.Branch
}

func (a *canoeBuild) GetSHA() string {
	if a == nil {
		return ""
	}
	return a.SHA
}

func (a *canoeBuild) GetShortSHA() string {
	if a == nil {
		return ""
	}
	if len(a.SHA) < 7 {
		return a.SHA
	}
	return a.SHA[0:7]
}

func (a *canoeBuild) GetURL() string {
	if a == nil {
		return ""
	}
	return a.ArtifactURL
}

func (a *canoeBuild) GetRepoURL() string {
	if a == nil {
		return ""
	}
	return a.RepoURL
}

func (a *canoeBuild) GetCreated() time.Time {
	if a == nil {
		return time.Unix(0, 0)
	}
	return time.Now()
}

type parsedImg struct {
	host       string
	registryID string
	repo       string
	tag        string
}

// parseImage parses a ecs.ContainerDefinition string Image.
func parseImage(img string) (*parsedImg, error) {
	u, err := url.Parse("docker://" + img)
	if err != nil {
		return nil, err
	}
	host := strings.Split(u.Host, ".")
	path := strings.Split(u.Path, ":")
	return &parsedImg{
		host:       u.Host,
		registryID: host[0],
		repo:       path[0][1:],
		tag:        path[1],
	}, nil
}
