// Code generated by go-swagger; DO NOT EDIT.

package tenants

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudentity/acp-client-go/models"
)

// PatchConfigurationReader is a Reader for the PatchConfiguration structure.
type PatchConfigurationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchConfigurationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPatchConfigurationNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchConfigurationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPatchConfigurationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchConfigurationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchConfigurationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewPatchConfigurationUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchConfigurationNoContent creates a PatchConfigurationNoContent with default headers values
func NewPatchConfigurationNoContent() *PatchConfigurationNoContent {
	return &PatchConfigurationNoContent{}
}

/* PatchConfigurationNoContent describes a response with status code 204, with default header values.

patch applied
*/
type PatchConfigurationNoContent struct {
}

func (o *PatchConfigurationNoContent) Error() string {
	return fmt.Sprintf("[PATCH /api/system/configuration][%d] patchConfigurationNoContent ", 204)
}

func (o *PatchConfigurationNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchConfigurationBadRequest creates a PatchConfigurationBadRequest with default headers values
func NewPatchConfigurationBadRequest() *PatchConfigurationBadRequest {
	return &PatchConfigurationBadRequest{}
}

/* PatchConfigurationBadRequest describes a response with status code 400, with default header values.

HttpError
*/
type PatchConfigurationBadRequest struct {
	Payload *models.Error
}

func (o *PatchConfigurationBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/system/configuration][%d] patchConfigurationBadRequest  %+v", 400, o.Payload)
}
func (o *PatchConfigurationBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PatchConfigurationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConfigurationUnauthorized creates a PatchConfigurationUnauthorized with default headers values
func NewPatchConfigurationUnauthorized() *PatchConfigurationUnauthorized {
	return &PatchConfigurationUnauthorized{}
}

/* PatchConfigurationUnauthorized describes a response with status code 401, with default header values.

HttpError
*/
type PatchConfigurationUnauthorized struct {
	Payload *models.Error
}

func (o *PatchConfigurationUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/system/configuration][%d] patchConfigurationUnauthorized  %+v", 401, o.Payload)
}
func (o *PatchConfigurationUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PatchConfigurationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConfigurationForbidden creates a PatchConfigurationForbidden with default headers values
func NewPatchConfigurationForbidden() *PatchConfigurationForbidden {
	return &PatchConfigurationForbidden{}
}

/* PatchConfigurationForbidden describes a response with status code 403, with default header values.

HttpError
*/
type PatchConfigurationForbidden struct {
	Payload *models.Error
}

func (o *PatchConfigurationForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/system/configuration][%d] patchConfigurationForbidden  %+v", 403, o.Payload)
}
func (o *PatchConfigurationForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PatchConfigurationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConfigurationNotFound creates a PatchConfigurationNotFound with default headers values
func NewPatchConfigurationNotFound() *PatchConfigurationNotFound {
	return &PatchConfigurationNotFound{}
}

/* PatchConfigurationNotFound describes a response with status code 404, with default header values.

HttpError
*/
type PatchConfigurationNotFound struct {
	Payload *models.Error
}

func (o *PatchConfigurationNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/system/configuration][%d] patchConfigurationNotFound  %+v", 404, o.Payload)
}
func (o *PatchConfigurationNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PatchConfigurationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConfigurationUnprocessableEntity creates a PatchConfigurationUnprocessableEntity with default headers values
func NewPatchConfigurationUnprocessableEntity() *PatchConfigurationUnprocessableEntity {
	return &PatchConfigurationUnprocessableEntity{}
}

/* PatchConfigurationUnprocessableEntity describes a response with status code 422, with default header values.

HttpError
*/
type PatchConfigurationUnprocessableEntity struct {
	Payload *models.Error
}

func (o *PatchConfigurationUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PATCH /api/system/configuration][%d] patchConfigurationUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *PatchConfigurationUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *PatchConfigurationUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
