// Copyright 2014 Mark Wolfe. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package buildkite

import "fmt"

// OrganizationsService handles communication with the organization related
// methods of the buildkite API.
//
// buildkite API docs: https://buildkite.com/docs/api/organizations
type OrganizationsService struct {
	client *Client
}

// Organization represents a buildkite organization.
type Organization struct {
	ID          *string    `json:"id,omitempty"`
	URL         *string    `json:"url,omitempty"`
	Name        *string    `json:"name,omitempty"`
	Slug        *string    `json:"slug,omitempty"`
	Repository  *string    `json:"repository,omitempty"`
	ProjectsURL *string    `json:"projects_url,omitempty"`
	AgentsURL   *string    `json:"agents_url,omitempty"`
	CreatedAt   *Timestamp `json:"created_at,omitempty"`
}

// OrganizationListOptions specifies the optional parameters to the
// OrganizationsService.List method.
type OrganizationListOptions struct {
	ListOptions
}

// List the organizations for the current user.
//
// buildkite API docs: https://buildkite.com/docs/api/organizations#list-organizations
func (os *OrganizationsService) List(opt *OrganizationListOptions) ([]Organization, *Response, error) {
	var u string

	u = fmt.Sprintf("v1/organizations")

	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := os.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	orgs := new([]Organization)
	resp, err := os.client.Do(req, orgs)
	if err != nil {
		return nil, resp, err
	}

	return *orgs, resp, err
}
