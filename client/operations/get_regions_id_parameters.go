// Code generated by go-swagger; DO NOT EDIT.

package operations

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
)

// NewGetRegionsIDParams creates a new GetRegionsIDParams object
// with the default values initialized.
func NewGetRegionsIDParams() *GetRegionsIDParams {
	var ()
	return &GetRegionsIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRegionsIDParamsWithTimeout creates a new GetRegionsIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRegionsIDParamsWithTimeout(timeout time.Duration) *GetRegionsIDParams {
	var ()
	return &GetRegionsIDParams{

		timeout: timeout,
	}
}

// NewGetRegionsIDParamsWithContext creates a new GetRegionsIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRegionsIDParamsWithContext(ctx context.Context) *GetRegionsIDParams {
	var ()
	return &GetRegionsIDParams{

		Context: ctx,
	}
}

// NewGetRegionsIDParamsWithHTTPClient creates a new GetRegionsIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRegionsIDParamsWithHTTPClient(client *http.Client) *GetRegionsIDParams {
	var ()
	return &GetRegionsIDParams{
		HTTPClient: client,
	}
}

/*GetRegionsIDParams contains all the parameters to send to the API endpoint
for the get regions ID operation typically these are written to a http.Request
*/
type GetRegionsIDParams struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get regions ID params
func (o *GetRegionsIDParams) WithTimeout(timeout time.Duration) *GetRegionsIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get regions ID params
func (o *GetRegionsIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get regions ID params
func (o *GetRegionsIDParams) WithContext(ctx context.Context) *GetRegionsIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get regions ID params
func (o *GetRegionsIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get regions ID params
func (o *GetRegionsIDParams) WithHTTPClient(client *http.Client) *GetRegionsIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get regions ID params
func (o *GetRegionsIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get regions ID params
func (o *GetRegionsIDParams) WithID(id string) *GetRegionsIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get regions ID params
func (o *GetRegionsIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetRegionsIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
