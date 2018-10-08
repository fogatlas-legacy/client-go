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

// GetExternalendpointsIDReader is a Reader for the GetExternalendpointsID structure.
type GetExternalendpointsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetExternalendpointsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetExternalendpointsIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetExternalendpointsIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetExternalendpointsIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetExternalendpointsIDOK creates a GetExternalendpointsIDOK with default headers values
func NewGetExternalendpointsIDOK() *GetExternalendpointsIDOK {
	return &GetExternalendpointsIDOK{}
}

/*GetExternalendpointsIDOK handles this case with default header values.

Successfully returned the external endpoint
*/
type GetExternalendpointsIDOK struct {
	Payload *models.ExternalEndpoint
}

func (o *GetExternalendpointsIDOK) Error() string {
	return fmt.Sprintf("[GET /externalendpoints/{id}][%d] getExternalendpointsIdOK  %+v", 200, o.Payload)
}

func (o *GetExternalendpointsIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ExternalEndpoint)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalendpointsIDBadRequest creates a GetExternalendpointsIDBadRequest with default headers values
func NewGetExternalendpointsIDBadRequest() *GetExternalendpointsIDBadRequest {
	return &GetExternalendpointsIDBadRequest{}
}

/*GetExternalendpointsIDBadRequest handles this case with default header values.

Invalid request
*/
type GetExternalendpointsIDBadRequest struct {
	Payload *models.Message
}

func (o *GetExternalendpointsIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /externalendpoints/{id}][%d] getExternalendpointsIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetExternalendpointsIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalendpointsIDInternalServerError creates a GetExternalendpointsIDInternalServerError with default headers values
func NewGetExternalendpointsIDInternalServerError() *GetExternalendpointsIDInternalServerError {
	return &GetExternalendpointsIDInternalServerError{}
}

/*GetExternalendpointsIDInternalServerError handles this case with default header values.

Generic error
*/
type GetExternalendpointsIDInternalServerError struct {
	Payload *models.Message
}

func (o *GetExternalendpointsIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /externalendpoints/{id}][%d] getExternalendpointsIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetExternalendpointsIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
