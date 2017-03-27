package canoe

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"git.dev.pardot.com/Pardot/infrastructure/bread/generated/swagger/models"
)

// NewPhoneAuthenticationParams creates a new PhoneAuthenticationParams object
// with the default values initialized.
func NewPhoneAuthenticationParams() *PhoneAuthenticationParams {
	var ()
	return &PhoneAuthenticationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPhoneAuthenticationParamsWithTimeout creates a new PhoneAuthenticationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPhoneAuthenticationParamsWithTimeout(timeout time.Duration) *PhoneAuthenticationParams {
	var ()
	return &PhoneAuthenticationParams{

		timeout: timeout,
	}
}

// NewPhoneAuthenticationParamsWithContext creates a new PhoneAuthenticationParams object
// with the default values initialized, and the ability to set a context for a request
func NewPhoneAuthenticationParamsWithContext(ctx context.Context) *PhoneAuthenticationParams {
	var ()
	return &PhoneAuthenticationParams{

		Context: ctx,
	}
}

// NewPhoneAuthenticationParamsWithHTTPClient creates a new PhoneAuthenticationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPhoneAuthenticationParamsWithHTTPClient(client *http.Client) *PhoneAuthenticationParams {
	var ()
	return &PhoneAuthenticationParams{
		HTTPClient: client,
	}
}

/*PhoneAuthenticationParams contains all the parameters to send to the API endpoint
for the phone authentication operation typically these are written to a http.Request
*/
type PhoneAuthenticationParams struct {

	/*Body*/
	Body *models.BreadPhoneAuthenticationRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the phone authentication params
func (o *PhoneAuthenticationParams) WithTimeout(timeout time.Duration) *PhoneAuthenticationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the phone authentication params
func (o *PhoneAuthenticationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the phone authentication params
func (o *PhoneAuthenticationParams) WithContext(ctx context.Context) *PhoneAuthenticationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the phone authentication params
func (o *PhoneAuthenticationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the phone authentication params
func (o *PhoneAuthenticationParams) WithHTTPClient(client *http.Client) *PhoneAuthenticationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the phone authentication params
func (o *PhoneAuthenticationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the phone authentication params
func (o *PhoneAuthenticationParams) WithBody(body *models.BreadPhoneAuthenticationRequest) *PhoneAuthenticationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the phone authentication params
func (o *PhoneAuthenticationParams) SetBody(body *models.BreadPhoneAuthenticationRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PhoneAuthenticationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.Body == nil {
		o.Body = new(models.BreadPhoneAuthenticationRequest)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}