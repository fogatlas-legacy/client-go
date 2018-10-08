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

// DeleteExternalendpointsIDReader is a Reader for the DeleteExternalendpointsID structure.
type DeleteExternalendpointsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteExternalendpointsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteExternalendpointsIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteExternalendpointsIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewDeleteExternalendpointsIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteExternalendpointsIDNoContent creates a DeleteExternalendpointsIDNoContent with default headers values
func NewDeleteExternalendpointsIDNoContent() *DeleteExternalendpointsIDNoContent {
	return &DeleteExternalendpointsIDNoContent{}
}

/*DeleteExternalendpointsIDNoContent handles this case with default header values.

External endpoint deleted
*/
type DeleteExternalendpointsIDNoContent struct {
}

func (o *DeleteExternalendpointsIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /externalendpoints/{id}][%d] deleteExternalendpointsIdNoContent ", 204)
}

func (o *DeleteExternalendpointsIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteExternalendpointsIDBadRequest creates a DeleteExternalendpointsIDBadRequest with default headers values
func NewDeleteExternalendpointsIDBadRequest() *DeleteExternalendpointsIDBadRequest {
	return &DeleteExternalendpointsIDBadRequest{}
}

/*DeleteExternalendpointsIDBadRequest handles this case with default header values.

Invalid request
*/
type DeleteExternalendpointsIDBadRequest struct {
	Payload *models.Message
}

func (o *DeleteExternalendpointsIDBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /externalendpoints/{id}][%d] deleteExternalendpointsIdBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteExternalendpointsIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteExternalendpointsIDInternalServerError creates a DeleteExternalendpointsIDInternalServerError with default headers values
func NewDeleteExternalendpointsIDInternalServerError() *DeleteExternalendpointsIDInternalServerError {
	return &DeleteExternalendpointsIDInternalServerError{}
}

/*DeleteExternalendpointsIDInternalServerError handles this case with default header values.

Generic error
*/
type DeleteExternalendpointsIDInternalServerError struct {
	Payload *models.Message
}

func (o *DeleteExternalendpointsIDInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /externalendpoints/{id}][%d] deleteExternalendpointsIdInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteExternalendpointsIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}