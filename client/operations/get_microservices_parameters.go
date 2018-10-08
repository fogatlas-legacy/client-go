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

// NewGetMicroservicesParams creates a new GetMicroservicesParams object
// with the default values initialized.
func NewGetMicroservicesParams() *GetMicroservicesParams {
	var ()
	return &GetMicroservicesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetMicroservicesParamsWithTimeout creates a new GetMicroservicesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetMicroservicesParamsWithTimeout(timeout time.Duration) *GetMicroservicesParams {
	var ()
	return &GetMicroservicesParams{

		timeout: timeout,
	}
}

// NewGetMicroservicesParamsWithContext creates a new GetMicroservicesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetMicroservicesParamsWithContext(ctx context.Context) *GetMicroservicesParams {
	var ()
	return &GetMicroservicesParams{

		Context: ctx,
	}
}

// NewGetMicroservicesParamsWithHTTPClient creates a new GetMicroservicesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetMicroservicesParamsWithHTTPClient(client *http.Client) *GetMicroservicesParams {
	var ()
	return &GetMicroservicesParams{
		HTTPClient: client,
	}
}

/*GetMicroservicesParams contains all the parameters to send to the API endpoint
for the get microservices operation typically these are written to a http.Request
*/
type GetMicroservicesParams struct {

	/*NodeID
	  Returns microservices on the given node

	*/
	NodeID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get microservices params
func (o *GetMicroservicesParams) WithTimeout(timeout time.Duration) *GetMicroservicesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get microservices params
func (o *GetMicroservicesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get microservices params
func (o *GetMicroservicesParams) WithContext(ctx context.Context) *GetMicroservicesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get microservices params
func (o *GetMicroservicesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get microservices params
func (o *GetMicroservicesParams) WithHTTPClient(client *http.Client) *GetMicroservicesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get microservices params
func (o *GetMicroservicesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNodeID adds the nodeID to the get microservices params
func (o *GetMicroservicesParams) WithNodeID(nodeID *string) *GetMicroservicesParams {
	o.SetNodeID(nodeID)
	return o
}

// SetNodeID adds the nodeId to the get microservices params
func (o *GetMicroservicesParams) SetNodeID(nodeID *string) {
	o.NodeID = nodeID
}

// WriteToRequest writes these params to a swagger request
func (o *GetMicroservicesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.NodeID != nil {

		// query param node_id
		var qrNodeID string
		if o.NodeID != nil {
			qrNodeID = *o.NodeID
		}
		qNodeID := qrNodeID
		if qNodeID != "" {
			if err := r.SetQueryParam("node_id", qNodeID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
