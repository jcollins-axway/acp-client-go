// Code generated by go-swagger; DO NOT EDIT.

package openbanking

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudentity/acp-client-go/models"
)

// ObbrDataAccessConsentIntrospectReader is a Reader for the ObbrDataAccessConsentIntrospect structure.
type ObbrDataAccessConsentIntrospectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ObbrDataAccessConsentIntrospectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewObbrDataAccessConsentIntrospectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewObbrDataAccessConsentIntrospectUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewObbrDataAccessConsentIntrospectNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewObbrDataAccessConsentIntrospectOK creates a ObbrDataAccessConsentIntrospectOK with default headers values
func NewObbrDataAccessConsentIntrospectOK() *ObbrDataAccessConsentIntrospectOK {
	return &ObbrDataAccessConsentIntrospectOK{}
}

/* ObbrDataAccessConsentIntrospectOK describes a response with status code 200, with default header values.

IntrospectOBBRDataAccessConsentResponse
*/
type ObbrDataAccessConsentIntrospectOK struct {
	Payload *models.IntrospectOBBRDataAccessConsentResponse
}

func (o *ObbrDataAccessConsentIntrospectOK) Error() string {
	return fmt.Sprintf("[POST /{tid}/{aid}/open-banking-brasil/open-banking/consents/v1/consents/introspect][%d] obbrDataAccessConsentIntrospectOK  %+v", 200, o.Payload)
}
func (o *ObbrDataAccessConsentIntrospectOK) GetPayload() *models.IntrospectOBBRDataAccessConsentResponse {
	return o.Payload
}

func (o *ObbrDataAccessConsentIntrospectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IntrospectOBBRDataAccessConsentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewObbrDataAccessConsentIntrospectUnauthorized creates a ObbrDataAccessConsentIntrospectUnauthorized with default headers values
func NewObbrDataAccessConsentIntrospectUnauthorized() *ObbrDataAccessConsentIntrospectUnauthorized {
	return &ObbrDataAccessConsentIntrospectUnauthorized{}
}

/* ObbrDataAccessConsentIntrospectUnauthorized describes a response with status code 401, with default header values.

genericError
*/
type ObbrDataAccessConsentIntrospectUnauthorized struct {
	Payload *models.GenericError
}

func (o *ObbrDataAccessConsentIntrospectUnauthorized) Error() string {
	return fmt.Sprintf("[POST /{tid}/{aid}/open-banking-brasil/open-banking/consents/v1/consents/introspect][%d] obbrDataAccessConsentIntrospectUnauthorized  %+v", 401, o.Payload)
}
func (o *ObbrDataAccessConsentIntrospectUnauthorized) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *ObbrDataAccessConsentIntrospectUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewObbrDataAccessConsentIntrospectNotFound creates a ObbrDataAccessConsentIntrospectNotFound with default headers values
func NewObbrDataAccessConsentIntrospectNotFound() *ObbrDataAccessConsentIntrospectNotFound {
	return &ObbrDataAccessConsentIntrospectNotFound{}
}

/* ObbrDataAccessConsentIntrospectNotFound describes a response with status code 404, with default header values.

genericError
*/
type ObbrDataAccessConsentIntrospectNotFound struct {
	Payload *models.GenericError
}

func (o *ObbrDataAccessConsentIntrospectNotFound) Error() string {
	return fmt.Sprintf("[POST /{tid}/{aid}/open-banking-brasil/open-banking/consents/v1/consents/introspect][%d] obbrDataAccessConsentIntrospectNotFound  %+v", 404, o.Payload)
}
func (o *ObbrDataAccessConsentIntrospectNotFound) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *ObbrDataAccessConsentIntrospectNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
