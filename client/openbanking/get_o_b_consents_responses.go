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

// GetOBConsentsReader is a Reader for the GetOBConsents structure.
type GetOBConsentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOBConsentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOBConsentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetOBConsentsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetOBConsentsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetOBConsentsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetOBConsentsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOBConsentsOK creates a GetOBConsentsOK with default headers values
func NewGetOBConsentsOK() *GetOBConsentsOK {
	return &GetOBConsentsOK{}
}

/* GetOBConsentsOK describes a response with status code 200, with default header values.

GenericConsents
*/
type GetOBConsentsOK struct {
	Payload *models.GenericConsents
}

func (o *GetOBConsentsOK) Error() string {
	return fmt.Sprintf("[GET /api/system/{tid}/servers/{aid}/open-banking/consents][%d] getOBConsentsOK  %+v", 200, o.Payload)
}
func (o *GetOBConsentsOK) GetPayload() *models.GenericConsents {
	return o.Payload
}

func (o *GetOBConsentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericConsents)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOBConsentsBadRequest creates a GetOBConsentsBadRequest with default headers values
func NewGetOBConsentsBadRequest() *GetOBConsentsBadRequest {
	return &GetOBConsentsBadRequest{}
}

/* GetOBConsentsBadRequest describes a response with status code 400, with default header values.

HttpError
*/
type GetOBConsentsBadRequest struct {
	Payload *models.Error
}

func (o *GetOBConsentsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/system/{tid}/servers/{aid}/open-banking/consents][%d] getOBConsentsBadRequest  %+v", 400, o.Payload)
}
func (o *GetOBConsentsBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetOBConsentsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOBConsentsUnauthorized creates a GetOBConsentsUnauthorized with default headers values
func NewGetOBConsentsUnauthorized() *GetOBConsentsUnauthorized {
	return &GetOBConsentsUnauthorized{}
}

/* GetOBConsentsUnauthorized describes a response with status code 401, with default header values.

HttpError
*/
type GetOBConsentsUnauthorized struct {
	Payload *models.Error
}

func (o *GetOBConsentsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/system/{tid}/servers/{aid}/open-banking/consents][%d] getOBConsentsUnauthorized  %+v", 401, o.Payload)
}
func (o *GetOBConsentsUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetOBConsentsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOBConsentsForbidden creates a GetOBConsentsForbidden with default headers values
func NewGetOBConsentsForbidden() *GetOBConsentsForbidden {
	return &GetOBConsentsForbidden{}
}

/* GetOBConsentsForbidden describes a response with status code 403, with default header values.

HttpError
*/
type GetOBConsentsForbidden struct {
	Payload *models.Error
}

func (o *GetOBConsentsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/system/{tid}/servers/{aid}/open-banking/consents][%d] getOBConsentsForbidden  %+v", 403, o.Payload)
}
func (o *GetOBConsentsForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetOBConsentsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOBConsentsNotFound creates a GetOBConsentsNotFound with default headers values
func NewGetOBConsentsNotFound() *GetOBConsentsNotFound {
	return &GetOBConsentsNotFound{}
}

/* GetOBConsentsNotFound describes a response with status code 404, with default header values.

HttpError
*/
type GetOBConsentsNotFound struct {
	Payload *models.Error
}

func (o *GetOBConsentsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/system/{tid}/servers/{aid}/open-banking/consents][%d] getOBConsentsNotFound  %+v", 404, o.Payload)
}
func (o *GetOBConsentsNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetOBConsentsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
