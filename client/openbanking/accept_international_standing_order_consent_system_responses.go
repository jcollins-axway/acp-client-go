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

// AcceptInternationalStandingOrderConsentSystemReader is a Reader for the AcceptInternationalStandingOrderConsentSystem structure.
type AcceptInternationalStandingOrderConsentSystemReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AcceptInternationalStandingOrderConsentSystemReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAcceptInternationalStandingOrderConsentSystemOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewAcceptInternationalStandingOrderConsentSystemUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAcceptInternationalStandingOrderConsentSystemForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAcceptInternationalStandingOrderConsentSystemNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAcceptInternationalStandingOrderConsentSystemOK creates a AcceptInternationalStandingOrderConsentSystemOK with default headers values
func NewAcceptInternationalStandingOrderConsentSystemOK() *AcceptInternationalStandingOrderConsentSystemOK {
	return &AcceptInternationalStandingOrderConsentSystemOK{}
}

/* AcceptInternationalStandingOrderConsentSystemOK describes a response with status code 200, with default header values.

ConsentAccepted
*/
type AcceptInternationalStandingOrderConsentSystemOK struct {
	Payload *models.ConsentAccepted
}

func (o *AcceptInternationalStandingOrderConsentSystemOK) Error() string {
	return fmt.Sprintf("[POST /api/system/{tid}/open-banking/international-standing-order-consent/{login}/accept][%d] acceptInternationalStandingOrderConsentSystemOK  %+v", 200, o.Payload)
}
func (o *AcceptInternationalStandingOrderConsentSystemOK) GetPayload() *models.ConsentAccepted {
	return o.Payload
}

func (o *AcceptInternationalStandingOrderConsentSystemOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConsentAccepted)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAcceptInternationalStandingOrderConsentSystemUnauthorized creates a AcceptInternationalStandingOrderConsentSystemUnauthorized with default headers values
func NewAcceptInternationalStandingOrderConsentSystemUnauthorized() *AcceptInternationalStandingOrderConsentSystemUnauthorized {
	return &AcceptInternationalStandingOrderConsentSystemUnauthorized{}
}

/* AcceptInternationalStandingOrderConsentSystemUnauthorized describes a response with status code 401, with default header values.

HttpError
*/
type AcceptInternationalStandingOrderConsentSystemUnauthorized struct {
	Payload *models.Error
}

func (o *AcceptInternationalStandingOrderConsentSystemUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/system/{tid}/open-banking/international-standing-order-consent/{login}/accept][%d] acceptInternationalStandingOrderConsentSystemUnauthorized  %+v", 401, o.Payload)
}
func (o *AcceptInternationalStandingOrderConsentSystemUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *AcceptInternationalStandingOrderConsentSystemUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAcceptInternationalStandingOrderConsentSystemForbidden creates a AcceptInternationalStandingOrderConsentSystemForbidden with default headers values
func NewAcceptInternationalStandingOrderConsentSystemForbidden() *AcceptInternationalStandingOrderConsentSystemForbidden {
	return &AcceptInternationalStandingOrderConsentSystemForbidden{}
}

/* AcceptInternationalStandingOrderConsentSystemForbidden describes a response with status code 403, with default header values.

HttpError
*/
type AcceptInternationalStandingOrderConsentSystemForbidden struct {
	Payload *models.Error
}

func (o *AcceptInternationalStandingOrderConsentSystemForbidden) Error() string {
	return fmt.Sprintf("[POST /api/system/{tid}/open-banking/international-standing-order-consent/{login}/accept][%d] acceptInternationalStandingOrderConsentSystemForbidden  %+v", 403, o.Payload)
}
func (o *AcceptInternationalStandingOrderConsentSystemForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *AcceptInternationalStandingOrderConsentSystemForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAcceptInternationalStandingOrderConsentSystemNotFound creates a AcceptInternationalStandingOrderConsentSystemNotFound with default headers values
func NewAcceptInternationalStandingOrderConsentSystemNotFound() *AcceptInternationalStandingOrderConsentSystemNotFound {
	return &AcceptInternationalStandingOrderConsentSystemNotFound{}
}

/* AcceptInternationalStandingOrderConsentSystemNotFound describes a response with status code 404, with default header values.

HttpError
*/
type AcceptInternationalStandingOrderConsentSystemNotFound struct {
	Payload *models.Error
}

func (o *AcceptInternationalStandingOrderConsentSystemNotFound) Error() string {
	return fmt.Sprintf("[POST /api/system/{tid}/open-banking/international-standing-order-consent/{login}/accept][%d] acceptInternationalStandingOrderConsentSystemNotFound  %+v", 404, o.Payload)
}
func (o *AcceptInternationalStandingOrderConsentSystemNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *AcceptInternationalStandingOrderConsentSystemNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
