// Code generated by go-swagger; DO NOT EDIT.

package system

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudentity/acp-client-go/models"
)

// ListSystemAPIsReader is a Reader for the ListSystemAPIs structure.
type ListSystemAPIsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListSystemAPIsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListSystemAPIsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListSystemAPIsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListSystemAPIsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListSystemAPIsOK creates a ListSystemAPIsOK with default headers values
func NewListSystemAPIsOK() *ListSystemAPIsOK {
	return &ListSystemAPIsOK{}
}

/* ListSystemAPIsOK describes a response with status code 200, with default header values.

ServerAPIs
*/
type ListSystemAPIsOK struct {
	Payload *models.ServerAPIs
}

func (o *ListSystemAPIsOK) Error() string {
	return fmt.Sprintf("[GET /api/admin/{tid}/system/apis][%d] listSystemAPIsOK  %+v", 200, o.Payload)
}
func (o *ListSystemAPIsOK) GetPayload() *models.ServerAPIs {
	return o.Payload
}

func (o *ListSystemAPIsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServerAPIs)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListSystemAPIsUnauthorized creates a ListSystemAPIsUnauthorized with default headers values
func NewListSystemAPIsUnauthorized() *ListSystemAPIsUnauthorized {
	return &ListSystemAPIsUnauthorized{}
}

/* ListSystemAPIsUnauthorized describes a response with status code 401, with default header values.

HttpError
*/
type ListSystemAPIsUnauthorized struct {
	Payload *models.Error
}

func (o *ListSystemAPIsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/admin/{tid}/system/apis][%d] listSystemAPIsUnauthorized  %+v", 401, o.Payload)
}
func (o *ListSystemAPIsUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListSystemAPIsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListSystemAPIsForbidden creates a ListSystemAPIsForbidden with default headers values
func NewListSystemAPIsForbidden() *ListSystemAPIsForbidden {
	return &ListSystemAPIsForbidden{}
}

/* ListSystemAPIsForbidden describes a response with status code 403, with default header values.

HttpError
*/
type ListSystemAPIsForbidden struct {
	Payload *models.Error
}

func (o *ListSystemAPIsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/admin/{tid}/system/apis][%d] listSystemAPIsForbidden  %+v", 403, o.Payload)
}
func (o *ListSystemAPIsForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListSystemAPIsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
