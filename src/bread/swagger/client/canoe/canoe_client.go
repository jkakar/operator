package canoe

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new canoe API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for canoe API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CompleteTerraformDeploy complete terraform deploy API
*/
func (a *Client) CompleteTerraformDeploy(params *CompleteTerraformDeployParams) (*CompleteTerraformDeployOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCompleteTerraformDeployParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CompleteTerraformDeploy",
		Method:             "POST",
		PathPattern:        "/api/grpc/complete_terraform_deploy",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CompleteTerraformDeployReader{formats: a.formats},
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CompleteTerraformDeployOK), nil

}

/*
CreateTerraformDeploy create terraform deploy API
*/
func (a *Client) CreateTerraformDeploy(params *CreateTerraformDeployParams) (*CreateTerraformDeployOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateTerraformDeployParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateTerraformDeploy",
		Method:             "POST",
		PathPattern:        "/api/grpc/create_terraform_deploy",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateTerraformDeployReader{formats: a.formats},
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateTerraformDeployOK), nil

}

/*
PhoneAuthentication phone authentication API
*/
func (a *Client) PhoneAuthentication(params *PhoneAuthenticationParams) (*PhoneAuthenticationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPhoneAuthenticationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PhoneAuthentication",
		Method:             "POST",
		PathPattern:        "/api/grpc/phone_authentication",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PhoneAuthenticationReader{formats: a.formats},
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PhoneAuthenticationOK), nil

}

/*
UnlockTerraformProject unlock terraform project API
*/
func (a *Client) UnlockTerraformProject(params *UnlockTerraformProjectParams) (*UnlockTerraformProjectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUnlockTerraformProjectParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UnlockTerraformProject",
		Method:             "POST",
		PathPattern:        "/api/grpc/unlock_terraform_project",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UnlockTerraformProjectReader{formats: a.formats},
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UnlockTerraformProjectOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}