package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// BreadUnlockTerraformProjectRequest bread unlock terraform project request
// swagger:model breadUnlockTerraformProjectRequest
type BreadUnlockTerraformProjectRequest struct {

	// project
	Project string `json:"project,omitempty"`

	// user email
	UserEmail string `json:"user_email,omitempty"`
}

// Validate validates this bread unlock terraform project request
func (m *BreadUnlockTerraformProjectRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}