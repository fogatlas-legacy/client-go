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

// PutDynamicnodesIDReader is a Reader for the PutDynamicnodesID structure.
type PutDynamicnodesIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutDynamicnodesIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPutDynamicnodesIDCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPutDynamicnodesIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPutDynamicnodesIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutDynamicnodesIDCreated creates a PutDynamicnodesIDCreated with default headers values
func NewPutDynamicnodesIDCreated() *PutDynamicnodesIDCreated {
	return &PutDynamicnodesIDCreated{}
}

/*PutDynamicnodesIDCreated handles this case with default header values.

Created
*/
type PutDynamicnodesIDCreated struct {
}

func (o *PutDynamicnodesIDCreated) Error() string {
	return fmt.Sprintf("[PUT /dynamicnodes/{id}][%d] putDynamicnodesIdCreated ", 201)
}

func (o *PutDynamicnodesIDCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutDynamicnodesIDBadRequest creates a PutDynamicnodesIDBadRequest with default headers values
func NewPutDynamicnodesIDBadRequest() *PutDynamicnodesIDBadRequest {
	return &PutDynamicnodesIDBadRequest{}
}

/*PutDynamicnodesIDBadRequest handles this case with default header values.

Invalid request
*/
type PutDynamicnodesIDBadRequest struct {
	Payload *models.Message
}

func (o *PutDynamicnodesIDBadRequest) Error() string {
	return fmt.Sprintf("[PUT /dynamicnodes/{id}][%d] putDynamicnodesIdBadRequest  %+v", 400, o.Payload)
}

func (o *PutDynamicnodesIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutDynamicnodesIDInternalServerError creates a PutDynamicnodesIDInternalServerError with default headers values
func NewPutDynamicnodesIDInternalServerError() *PutDynamicnodesIDInternalServerError {
	return &PutDynamicnodesIDInternalServerError{}
}

/*PutDynamicnodesIDInternalServerError handles this case with default header values.

Generic error
*/
type PutDynamicnodesIDInternalServerError struct {
	Payload *models.Message
}

func (o *PutDynamicnodesIDInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /dynamicnodes/{id}][%d] putDynamicnodesIdInternalServerError  %+v", 500, o.Payload)
}

func (o *PutDynamicnodesIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}