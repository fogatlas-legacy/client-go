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

	"github.com/fogatlas/client-go/models"
)

// NewPutDynamicnodesIDParams creates a new PutDynamicnodesIDParams object
// with the default values initialized.
func NewPutDynamicnodesIDParams() *PutDynamicnodesIDParams {
	var ()
	return &PutDynamicnodesIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutDynamicnodesIDParamsWithTimeout creates a new PutDynamicnodesIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutDynamicnodesIDParamsWithTimeout(timeout time.Duration) *PutDynamicnodesIDParams {
	var ()
	return &PutDynamicnodesIDParams{

		timeout: timeout,
	}
}

// NewPutDynamicnodesIDParamsWithContext creates a new PutDynamicnodesIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutDynamicnodesIDParamsWithContext(ctx context.Context) *PutDynamicnodesIDParams {
	var ()
	return &PutDynamicnodesIDParams{

		Context: ctx,
	}
}

// NewPutDynamicnodesIDParamsWithHTTPClient creates a new PutDynamicnodesIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutDynamicnodesIDParamsWithHTTPClient(client *http.Client) *PutDynamicnodesIDParams {
	var ()
	return &PutDynamicnodesIDParams{
		HTTPClient: client,
	}
}

/*PutDynamicnodesIDParams contains all the parameters to send to the API endpoint
for the put dynamicnodes ID operation typically these are written to a http.Request
*/
type PutDynamicnodesIDParams struct {

	/*Dynamicnode*/
	Dynamicnode *models.DynamicNode
	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put dynamicnodes ID params
func (o *PutDynamicnodesIDParams) WithTimeout(timeout time.Duration) *PutDynamicnodesIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put dynamicnodes ID params
func (o *PutDynamicnodesIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put dynamicnodes ID params
func (o *PutDynamicnodesIDParams) WithContext(ctx context.Context) *PutDynamicnodesIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put dynamicnodes ID params
func (o *PutDynamicnodesIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put dynamicnodes ID params
func (o *PutDynamicnodesIDParams) WithHTTPClient(client *http.Client) *PutDynamicnodesIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put dynamicnodes ID params
func (o *PutDynamicnodesIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDynamicnode adds the dynamicnode to the put dynamicnodes ID params
func (o *PutDynamicnodesIDParams) WithDynamicnode(dynamicnode *models.DynamicNode) *PutDynamicnodesIDParams {
	o.SetDynamicnode(dynamicnode)
	return o
}

// SetDynamicnode adds the dynamicnode to the put dynamicnodes ID params
func (o *PutDynamicnodesIDParams) SetDynamicnode(dynamicnode *models.DynamicNode) {
	o.Dynamicnode = dynamicnode
}

// WithID adds the id to the put dynamicnodes ID params
func (o *PutDynamicnodesIDParams) WithID(id string) *PutDynamicnodesIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the put dynamicnodes ID params
func (o *PutDynamicnodesIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PutDynamicnodesIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Dynamicnode != nil {
		if err := r.SetBodyParam(o.Dynamicnode); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
