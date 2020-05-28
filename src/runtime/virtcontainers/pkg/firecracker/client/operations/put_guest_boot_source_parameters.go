// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/kata-containers/kata-containers/src/runtime/virtcontainers/pkg/firecracker/client/models"
)

// NewPutGuestBootSourceParams creates a new PutGuestBootSourceParams object
// with the default values initialized.
func NewPutGuestBootSourceParams() *PutGuestBootSourceParams {
	var ()
	return &PutGuestBootSourceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutGuestBootSourceParamsWithTimeout creates a new PutGuestBootSourceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutGuestBootSourceParamsWithTimeout(timeout time.Duration) *PutGuestBootSourceParams {
	var ()
	return &PutGuestBootSourceParams{

		timeout: timeout,
	}
}

// NewPutGuestBootSourceParamsWithContext creates a new PutGuestBootSourceParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutGuestBootSourceParamsWithContext(ctx context.Context) *PutGuestBootSourceParams {
	var ()
	return &PutGuestBootSourceParams{

		Context: ctx,
	}
}

// NewPutGuestBootSourceParamsWithHTTPClient creates a new PutGuestBootSourceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutGuestBootSourceParamsWithHTTPClient(client *http.Client) *PutGuestBootSourceParams {
	var ()
	return &PutGuestBootSourceParams{
		HTTPClient: client,
	}
}

/*PutGuestBootSourceParams contains all the parameters to send to the API endpoint
for the put guest boot source operation typically these are written to a http.Request
*/
type PutGuestBootSourceParams struct {

	/*Body
	  Guest boot source properties

	*/
	Body *models.BootSource

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put guest boot source params
func (o *PutGuestBootSourceParams) WithTimeout(timeout time.Duration) *PutGuestBootSourceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put guest boot source params
func (o *PutGuestBootSourceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put guest boot source params
func (o *PutGuestBootSourceParams) WithContext(ctx context.Context) *PutGuestBootSourceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put guest boot source params
func (o *PutGuestBootSourceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put guest boot source params
func (o *PutGuestBootSourceParams) WithHTTPClient(client *http.Client) *PutGuestBootSourceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put guest boot source params
func (o *PutGuestBootSourceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put guest boot source params
func (o *PutGuestBootSourceParams) WithBody(body *models.BootSource) *PutGuestBootSourceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put guest boot source params
func (o *PutGuestBootSourceParams) SetBody(body *models.BootSource) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PutGuestBootSourceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}