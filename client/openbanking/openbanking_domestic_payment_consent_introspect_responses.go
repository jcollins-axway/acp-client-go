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

// OpenbankingDomesticPaymentConsentIntrospectReader is a Reader for the OpenbankingDomesticPaymentConsentIntrospect structure.
type OpenbankingDomesticPaymentConsentIntrospectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OpenbankingDomesticPaymentConsentIntrospectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOpenbankingDomesticPaymentConsentIntrospectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewOpenbankingDomesticPaymentConsentIntrospectUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewOpenbankingDomesticPaymentConsentIntrospectNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewOpenbankingDomesticPaymentConsentIntrospectOK creates a OpenbankingDomesticPaymentConsentIntrospectOK with default headers values
func NewOpenbankingDomesticPaymentConsentIntrospectOK() *OpenbankingDomesticPaymentConsentIntrospectOK {
	return &OpenbankingDomesticPaymentConsentIntrospectOK{}
}

/* OpenbankingDomesticPaymentConsentIntrospectOK describes a response with status code 200, with default header values.

IntrospectOpenbankingDomesticPaymentConsentResponse
*/
type OpenbankingDomesticPaymentConsentIntrospectOK struct {
	Payload *models.IntrospectOpenbankingDomesticPaymentConsentResponse
}

func (o *OpenbankingDomesticPaymentConsentIntrospectOK) Error() string {
	return fmt.Sprintf("[POST /{tid}/{aid}/open-banking/v3.1/pisp/domestic-payment-consents/introspect][%d] openbankingDomesticPaymentConsentIntrospectOK  %+v", 200, o.Payload)
}
func (o *OpenbankingDomesticPaymentConsentIntrospectOK) GetPayload() *models.IntrospectOpenbankingDomesticPaymentConsentResponse {
	return o.Payload
}

func (o *OpenbankingDomesticPaymentConsentIntrospectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IntrospectOpenbankingDomesticPaymentConsentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpenbankingDomesticPaymentConsentIntrospectUnauthorized creates a OpenbankingDomesticPaymentConsentIntrospectUnauthorized with default headers values
func NewOpenbankingDomesticPaymentConsentIntrospectUnauthorized() *OpenbankingDomesticPaymentConsentIntrospectUnauthorized {
	return &OpenbankingDomesticPaymentConsentIntrospectUnauthorized{}
}

/* OpenbankingDomesticPaymentConsentIntrospectUnauthorized describes a response with status code 401, with default header values.

genericError
*/
type OpenbankingDomesticPaymentConsentIntrospectUnauthorized struct {
	Payload *models.GenericError
}

func (o *OpenbankingDomesticPaymentConsentIntrospectUnauthorized) Error() string {
	return fmt.Sprintf("[POST /{tid}/{aid}/open-banking/v3.1/pisp/domestic-payment-consents/introspect][%d] openbankingDomesticPaymentConsentIntrospectUnauthorized  %+v", 401, o.Payload)
}
func (o *OpenbankingDomesticPaymentConsentIntrospectUnauthorized) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *OpenbankingDomesticPaymentConsentIntrospectUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpenbankingDomesticPaymentConsentIntrospectNotFound creates a OpenbankingDomesticPaymentConsentIntrospectNotFound with default headers values
func NewOpenbankingDomesticPaymentConsentIntrospectNotFound() *OpenbankingDomesticPaymentConsentIntrospectNotFound {
	return &OpenbankingDomesticPaymentConsentIntrospectNotFound{}
}

/* OpenbankingDomesticPaymentConsentIntrospectNotFound describes a response with status code 404, with default header values.

genericError
*/
type OpenbankingDomesticPaymentConsentIntrospectNotFound struct {
	Payload *models.GenericError
}

func (o *OpenbankingDomesticPaymentConsentIntrospectNotFound) Error() string {
	return fmt.Sprintf("[POST /{tid}/{aid}/open-banking/v3.1/pisp/domestic-payment-consents/introspect][%d] openbankingDomesticPaymentConsentIntrospectNotFound  %+v", 404, o.Payload)
}
func (o *OpenbankingDomesticPaymentConsentIntrospectNotFound) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *OpenbankingDomesticPaymentConsentIntrospectNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
