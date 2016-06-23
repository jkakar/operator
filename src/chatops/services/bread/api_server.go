package bread

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"text/tabwriter"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ecr"
	"github.com/aws/aws-sdk-go/service/ecs"
	"github.com/sr/operator"
	"golang.org/x/net/context"
)

const defaultPlan = "BREAD-BREAD"

type apiServer struct {
	ecs       *ecs.ECS
	ecr       *ecr.ECR
	bamboo    *http.Client
	bambooURL *url.URL
	apps      map[string]string
	ecsSvc    string
}

type parsedImg struct {
	host       string
	registryID string
	repo       string
	tag        string
}

func newAPIServer(
	ecs *ecs.ECS,
	ecr *ecr.ECR,
	bamboo *http.Client,
	bambooURL *url.URL,
	apps map[string]string,
	ecsSvc string,
	_ int,
) (*apiServer, error) {
	return &apiServer{
		ecs,
		ecr,
		bamboo,
		bambooURL,
		apps,
		ecsSvc,
	}, nil
}

func (s *apiServer) ListApps(ctx context.Context, in *ListAppsRequest) (*ListAppsResponse, error) {
	apps := make([]string, len(s.apps))
	i := 0
	for _, s := range s.apps {
		apps[i] = s
		i = i + 1
	}
	return &ListAppsResponse{
		Output: &operator.Output{
			PlainText: fmt.Sprintf("deployable apps: %s", strings.Join(apps, ", ")),
		},
	}, nil
}

func (s *apiServer) ListBuilds(ctx context.Context, in *ListBuildsRequest) (*ListBuildsResponse, error) {
	var plan string
	if in.Plan == "" {
		plan = defaultPlan
	} else {
		plan = in.Plan
	}
	resp, err := s.bamboo.Get(fmt.Sprintf("%s/rest/api/latest/result/%s", s.bambooURL, plan))
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("bamboo request failed with status %d", resp.StatusCode)
	}
	var data struct {
		Results struct {
			Result []struct {
				LifeCycleState string `json:"lifeCycleState"`
				Key            string `json:"key"`
			} `json:"result"`
		} `json:"results"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	var (
		out bytes.Buffer
		w   tabwriter.Writer
	)
	w.Init(&out, 0, 4, 1, '\t', 0)
	fmt.Fprintf(&w, "%s\t%s\t%s\n", "ID", "STATUS", "URL")
	for _, build := range data.Results.Result {
		fmt.Fprintf(
			&w,
			"%s\t%s\t%s\n",
			build.Key,
			build.LifeCycleState,
			fmt.Sprintf("%s/%s", s.bambooURL, build.Key),
		)
	}
	if err := w.Flush(); err != nil {
		return nil, err
	}
	return &ListBuildsResponse{
		Output: &operator.Output{
			PlainText: out.String(),
		},
	}, nil
}

func (s *apiServer) EcsDeploy(ctx context.Context, in *EcsDeployRequest) (*EcsDeployResponse, error) {
	var cluster string
	cluster, ok := s.apps[in.App]
	if !ok {
		return nil, fmt.Errorf("no such app: %s", in.App)
	}
	svc, err := s.ecs.DescribeServices(
		&ecs.DescribeServicesInput{
			Services: []*string{aws.String(s.ecsSvc)},
			Cluster:  aws.String(cluster),
		},
	)
	if err != nil {
		return nil, err
	}
	if len(svc.Services) != 1 {
		return nil, errors.New("bogus response")
	}
	out, err := s.ecs.DescribeTaskDefinition(
		&ecs.DescribeTaskDefinitionInput{
			TaskDefinition: svc.Services[0].TaskDefinition,
		},
	)
	if err != nil {
		return nil, err
	}
	curImg, err := parseImage(*out.TaskDefinition.ContainerDefinitions[0].Image)
	if curImg.tag == in.Build {
		return nil, fmt.Errorf("build %s already deployed", in.Build)
	}
	if err != nil {
		return nil, err
	}
	img := fmt.Sprintf("%s/%s:%s", curImg.host, curImg.repo, in.Build)
	var nextToken *string
	found := false
OuterLoop:
	for {
		images, err := s.ecr.ListImages(
			&ecr.ListImagesInput{
				MaxResults:     aws.Int64(100),
				NextToken:      nextToken,
				RegistryId:     aws.String(curImg.registryID),
				RepositoryName: aws.String(curImg.repo),
			},
		)
		if err != nil {
			return nil, err
		}
		nextToken = images.NextToken
		if err != nil || len(images.ImageIds) == 0 {
			break OuterLoop
		}
		for _, i := range images.ImageIds {
			if i.ImageTag == nil {
				continue
			}
			if i.ImageTag != nil && *i.ImageTag == in.Build {
				found = true
				break OuterLoop
			}
		}

	}
	if !found {
		return nil, fmt.Errorf("image for build %s not found in Docker repository", in.Build)
	}
	out.TaskDefinition.ContainerDefinitions[0].Image = aws.String(img)
	task, err := s.ecs.RegisterTaskDefinition(
		&ecs.RegisterTaskDefinitionInput{
			ContainerDefinitions: out.TaskDefinition.ContainerDefinitions,
			Family:               out.TaskDefinition.Family,
			Volumes:              out.TaskDefinition.Volumes,
		},
	)
	if err != nil {
		return nil, err
	}
	_, err = s.ecs.UpdateService(
		&ecs.UpdateServiceInput{
			Cluster:        svc.Services[0].ClusterArn,
			Service:        svc.Services[0].ServiceName,
			TaskDefinition: task.TaskDefinition.TaskDefinitionArn,
		},
	)
	if err != nil {
		return nil, err
	}
	return &EcsDeployResponse{
		Output: &operator.Output{
			PlainText: fmt.Sprintf("deployed %s to %s", in.App, in.Build),
		},
	}, nil
}

// parseImage parses a ecs.ContainerDefinition string Image.
func parseImage(img string) (*parsedImg, error) {
	u, err := url.Parse("ecr://" + img)
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
