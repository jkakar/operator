package bread

import (
	"bytes"
	"errors"
	"fmt"

	"git.dev.pardot.com/Pardot/bread/jira"
	"github.com/sr/operator"
	"golang.org/x/net/context"

	"git.dev.pardot.com/Pardot/bread/pb"
)

type ticketsServer struct {
	operator.Sender
	jira    jira.Client
	project string
}

func NewTicketsServer(sender operator.Sender, jira jira.Client, project string) (breadpb.TicketsServer, error) {
	if jira == nil {
		return nil, errors.New("requirement argument 'jira' is nil")
	}
	if project == "" {
		return nil, errors.New("requirement argument 'project' is nil")
	}
	return &ticketsServer{sender, jira, project}, nil
}

func (s *ticketsServer) Mine(ctx context.Context, req *breadpb.MyIssuesRequest) (*operator.Response, error) {
	email := operator.GetUserEmail(req)
	if email == "" {
		return nil, errors.New("unable to retrieve list of assigned tickets without and email address")
	}

	tickets, err := s.jira.Search(
		ctx,
		fmt.Sprintf("project = '%s' AND Resolution IS NULL AND assignee = '%s'", s.project, email),
		[]string{"key", "summary", "status"},
		100,
	)
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve tickets: %s", err)
	}
	var txt, html bytes.Buffer
	_, _ = html.WriteString("<table><tr><th>Issue ID</th><th>Status</th><th>Summary</th></tr>")
	for _, issue := range tickets {
		if issue.Fields == nil {
			continue
		}
		fmt.Fprintf(&txt, "%s %s %s\n", issue.Key, issue.StatusName(), issue.Fields.Summary)
		fmt.Fprintf(
			&html,
			"<tr><td>%s</td><td>%s</td><td>%s</td></tr>\n",
			fmt.Sprintf(`<a href="%s">%s</a>`, issue.HTMLURL(), issue.Key),
			issue.StatusName(),
			issue.Fields.Summary,
		)
	}
	return operator.Reply(ctx, s, req, &operator.Message{Text: txt.String(), HTML: html.String()})
}