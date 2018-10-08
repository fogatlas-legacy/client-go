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

// PutNodesIDReader is a Reader for the PutNodesID structure.
type PutNodesIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutNodesIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPutNodesIDCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPutNodesIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPutNodesIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutNodesIDCreated creates a PutNodesIDCreated with default headers values
func NewPutNodesIDCreated() *PutNodesIDCreated {
	return &PutNodesIDCreated{}
}

/*PutNodesIDCreated handles this case with default header values.

Created
*/
type PutNodesIDCreated struct {
}

func (o *PutNodesIDCreated) Error() string {
	return fmt.Sprintf("[PUT /nodes/{id}][%d] putNodesIdCreated ", 201)
}

func (o *PutNodesIDCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutNodesIDBadRequest creates a PutNodesIDBadRequest with default headers values
func NewPutNodesIDBadRequest() *PutNodesIDBadRequest {
	return &PutNodesIDBadRequest{}
}

/*PutNodesIDBadRequest handles this case with default header values.

Invalid request
*/
type PutNodesIDBadRequest struct {
	Payload *models.Message
}

func (o *PutNodesIDBadRequest) Error() string {
	return fmt.Sprintf("[PUT /nodes/{id}][%d] putNodesIdBadRequest  %+v", 400, o.Payload)
}

func (o *PutNodesIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutNodesIDInternalServerError creates a PutNodesIDInternalServerError with default headers values
func NewPutNodesIDInternalServerError() *PutNodesIDInternalServerError {
	return &PutNodesIDInternalServerError{}
}

/*PutNodesIDInternalServerError handles this case with default header values.

Generic error
*/
type PutNodesIDInternalServerError struct {
	Payload *models.Message
}

func (o *PutNodesIDInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /nodes/{id}][%d] putNodesIdInternalServerError  %+v", 500, o.Payload)
}

func (o *PutNodesIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
