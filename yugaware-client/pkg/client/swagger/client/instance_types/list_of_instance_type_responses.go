// Code generated by go-swagger; DO NOT EDIT.

package instance_types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/yugabyte/yb-tools/yugaware-client/pkg/client/swagger/models"
)

// ListOfInstanceTypeReader is a Reader for the ListOfInstanceType structure.
type ListOfInstanceTypeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListOfInstanceTypeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListOfInstanceTypeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewListOfInstanceTypeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListOfInstanceTypeOK creates a ListOfInstanceTypeOK with default headers values
func NewListOfInstanceTypeOK() *ListOfInstanceTypeOK {
	return &ListOfInstanceTypeOK{}
}

/* ListOfInstanceTypeOK describes a response with status code 200, with default header values.

successful operation
*/
type ListOfInstanceTypeOK struct {
	Payload []models.PlatformResults
}

func (o *ListOfInstanceTypeOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/customers/{cUUID}/providers/{pUUID}/instance_types][%d] listOfInstanceTypeOK  %+v", 200, o.Payload)
}
func (o *ListOfInstanceTypeOK) GetPayload() []models.PlatformResults {
	return o.Payload
}

func (o *ListOfInstanceTypeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListOfInstanceTypeInternalServerError creates a ListOfInstanceTypeInternalServerError with default headers values
func NewListOfInstanceTypeInternalServerError() *ListOfInstanceTypeInternalServerError {
	return &ListOfInstanceTypeInternalServerError{}
}

/* ListOfInstanceTypeInternalServerError describes a response with status code 500, with default header values.

If there was a server or database issue when listing the instance types
*/
type ListOfInstanceTypeInternalServerError struct {
	Payload *models.YBPError
}

func (o *ListOfInstanceTypeInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/customers/{cUUID}/providers/{pUUID}/instance_types][%d] listOfInstanceTypeInternalServerError  %+v", 500, o.Payload)
}
func (o *ListOfInstanceTypeInternalServerError) GetPayload() *models.YBPError {
	return o.Payload
}

func (o *ListOfInstanceTypeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.YBPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
