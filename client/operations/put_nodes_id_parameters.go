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

// NewPutNodesIDParams creates a new PutNodesIDParams object
// with the default values initialized.
func NewPutNodesIDParams() *PutNodesIDParams {
	var ()
	return &PutNodesIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutNodesIDParamsWithTimeout creates a new PutNodesIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutNodesIDParamsWithTimeout(timeout time.Duration) *PutNodesIDParams {
	var ()
	return &PutNodesIDParams{

		timeout: timeout,
	}
}

// NewPutNodesIDParamsWithContext creates a new PutNodesIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutNodesIDParamsWithContext(ctx context.Context) *PutNodesIDParams {
	var ()
	return &PutNodesIDParams{

		Context: ctx,
	}
}

// NewPutNodesIDParamsWithHTTPClient creates a new PutNodesIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutNodesIDParamsWithHTTPClient(client *http.Client) *PutNodesIDParams {
	var ()
	return &PutNodesIDParams{
		HTTPClient: client,
	}
}

/*PutNodesIDParams contains all the parameters to send to the API endpoint
for the put nodes ID operation typically these are written to a http.Request
*/
type PutNodesIDParams struct {

	/*ID*/
	ID string
	/*Node*/
	Node *models.Node

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put nodes ID params
func (o *PutNodesIDParams) WithTimeout(timeout time.Duration) *PutNodesIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put nodes ID params
func (o *PutNodesIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put nodes ID params
func (o *PutNodesIDParams) WithContext(ctx context.Context) *PutNodesIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put nodes ID params
func (o *PutNodesIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put nodes ID params
func (o *PutNodesIDParams) WithHTTPClient(client *http.Client) *PutNodesIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put nodes ID params
func (o *PutNodesIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the put nodes ID params
func (o *PutNodesIDParams) WithID(id string) *PutNodesIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the put nodes ID params
func (o *PutNodesIDParams) SetID(id string) {
	o.ID = id
}

// WithNode adds the node to the put nodes ID params
func (o *PutNodesIDParams) WithNode(node *models.Node) *PutNodesIDParams {
	o.SetNode(node)
	return o
}

// SetNode adds the node to the put nodes ID params
func (o *PutNodesIDParams) SetNode(node *models.Node) {
	o.Node = node
}

// WriteToRequest writes these params to a swagger request
func (o *PutNodesIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if o.Node != nil {
		if err := r.SetBodyParam(o.Node); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
