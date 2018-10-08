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

// PatchDeploymentsNameReader is a Reader for the PatchDeploymentsName structure.
type PatchDeploymentsNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchDeploymentsNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPatchDeploymentsNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPatchDeploymentsNameBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPatchDeploymentsNameInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchDeploymentsNameOK creates a PatchDeploymentsNameOK with default headers values
func NewPatchDeploymentsNameOK() *PatchDeploymentsNameOK {
	return &PatchDeploymentsNameOK{}
}

/*PatchDeploymentsNameOK handles this case with default header values.

Updated
*/
type PatchDeploymentsNameOK struct {
}

func (o *PatchDeploymentsNameOK) Error() string {
	return fmt.Sprintf("[PATCH /deployments/{name}][%d] patchDeploymentsNameOK ", 200)
}

func (o *PatchDeploymentsNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchDeploymentsNameBadRequest creates a PatchDeploymentsNameBadRequest with default headers values
func NewPatchDeploymentsNameBadRequest() *PatchDeploymentsNameBadRequest {
	return &PatchDeploymentsNameBadRequest{}
}

/*PatchDeploymentsNameBadRequest handles this case with default header values.

Invalid request
*/
type PatchDeploymentsNameBadRequest struct {
	Payload *models.Message
}

func (o *PatchDeploymentsNameBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /deployments/{name}][%d] patchDeploymentsNameBadRequest  %+v", 400, o.Payload)
}

func (o *PatchDeploymentsNameBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchDeploymentsNameInternalServerError creates a PatchDeploymentsNameInternalServerError with default headers values
func NewPatchDeploymentsNameInternalServerError() *PatchDeploymentsNameInternalServerError {
	return &PatchDeploymentsNameInternalServerError{}
}

/*PatchDeploymentsNameInternalServerError handles this case with default header values.

Generic error
*/
type PatchDeploymentsNameInternalServerError struct {
	Payload *models.Message
}

func (o *PatchDeploymentsNameInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /deployments/{name}][%d] patchDeploymentsNameInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchDeploymentsNameInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
