package canoe

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"bread/swagger/models"
)

// NewCreateTerraformDeployParams creates a new CreateTerraformDeployParams object
// with the default values initialized.
func NewCreateTerraformDeployParams() *CreateTerraformDeployParams {
	var ()
	return &CreateTerraformDeployParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateTerraformDeployParamsWithTimeout creates a new CreateTerraformDeployParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateTerraformDeployParamsWithTimeout(timeout time.Duration) *CreateTerraformDeployParams {
	var ()
	return &CreateTerraformDeployParams{

		timeout: timeout,
	}
}

// NewCreateTerraformDeployParamsWithContext creates a new CreateTerraformDeployParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateTerraformDeployParamsWithContext(ctx context.Context) *CreateTerraformDeployParams {
	var ()
	return &CreateTerraformDeployParams{

		Context: ctx,
	}
}

/*CreateTerraformDeployParams contains all the parameters to send to the API endpoint
for the create terraform deploy operation typically these are written to a http.Request
*/
type CreateTerraformDeployParams struct {

	/*Body*/
	Body *models.CanoeCreateTerraformDeployRequest

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the create terraform deploy params
func (o *CreateTerraformDeployParams) WithTimeout(timeout time.Duration) *CreateTerraformDeployParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create terraform deploy params
func (o *CreateTerraformDeployParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create terraform deploy params
func (o *CreateTerraformDeployParams) WithContext(ctx context.Context) *CreateTerraformDeployParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create terraform deploy params
func (o *CreateTerraformDeployParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the create terraform deploy params
func (o *CreateTerraformDeployParams) WithBody(body *models.CanoeCreateTerraformDeployRequest) *CreateTerraformDeployParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create terraform deploy params
func (o *CreateTerraformDeployParams) SetBody(body *models.CanoeCreateTerraformDeployRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateTerraformDeployParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.Body == nil {
		o.Body = new(models.CanoeCreateTerraformDeployRequest)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
