// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/fogatlas/client-go/models"
)

// GetNodesIDReader is a Reader for the GetNodesID structure.
type GetNodesIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNodesIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetNodesIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetNodesIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetNodesIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetNodesIDOK creates a GetNodesIDOK with default headers values
func NewGetNodesIDOK() *GetNodesIDOK {
	return &GetNodesIDOK{}
}

/*GetNodesIDOK handles this case with default header values.

Successfully returned the node
*/
type GetNodesIDOK struct {
	Payload *models.Node
}

func (o *GetNodesIDOK) Error() string {
	return fmt.Sprintf("[GET /nodes/{id}][%d] getNodesIdOK  %+v", 200, o.Payload)
}

func (o *GetNodesIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Node)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNodesIDBadRequest creates a GetNodesIDBadRequest with default headers values
func NewGetNodesIDBadRequest() *GetNodesIDBadRequest {
	return &GetNodesIDBadRequest{}
}

/*GetNodesIDBadRequest handles this case with default header values.

Invalid request
*/
type GetNodesIDBadRequest struct {
	Payload *models.Message
}

func (o *GetNodesIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /nodes/{id}][%d] getNodesIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetNodesIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNodesIDInternalServerError creates a GetNodesIDInternalServerError with default headers values
func NewGetNodesIDInternalServerError() *GetNodesIDInternalServerError {
	return &GetNodesIDInternalServerError{}
}

/*GetNodesIDInternalServerError handles this case with default header values.

Generic error
*/
type GetNodesIDInternalServerError struct {
	Payload *models.Message
}

func (o *GetNodesIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /nodes/{id}][%d] getNodesIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetNodesIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
