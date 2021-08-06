// Code generated by go-swagger; DO NOT EDIT.

package scopes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudentity/acp-client-go/models"
)

// GetScopeReader is a Reader for the GetScope structure.
type GetScopeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetScopeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetScopeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetScopeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetScopeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetScopeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetScopeOK creates a GetScopeOK with default headers values
func NewGetScopeOK() *GetScopeOK {
	return &GetScopeOK{}
}

/* GetScopeOK describes a response with status code 200, with default header values.

ScopeWithService
*/
type GetScopeOK struct {
	Payload *models.ScopeWithService
}

func (o *GetScopeOK) Error() string {
	return fmt.Sprintf("[GET /api/admin/{tid}/scopes/{scp}][%d] getScopeOK  %+v", 200, o.Payload)
}
func (o *GetScopeOK) GetPayload() *models.ScopeWithService {
	return o.Payload
}

func (o *GetScopeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ScopeWithService)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScopeUnauthorized creates a GetScopeUnauthorized with default headers values
func NewGetScopeUnauthorized() *GetScopeUnauthorized {
	return &GetScopeUnauthorized{}
}

/* GetScopeUnauthorized describes a response with status code 401, with default header values.

HttpError
*/
type GetScopeUnauthorized struct {
	Payload *models.Error
}

func (o *GetScopeUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/admin/{tid}/scopes/{scp}][%d] getScopeUnauthorized  %+v", 401, o.Payload)
}
func (o *GetScopeUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetScopeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScopeForbidden creates a GetScopeForbidden with default headers values
func NewGetScopeForbidden() *GetScopeForbidden {
	return &GetScopeForbidden{}
}

/* GetScopeForbidden describes a response with status code 403, with default header values.

HttpError
*/
type GetScopeForbidden struct {
	Payload *models.Error
}

func (o *GetScopeForbidden) Error() string {
	return fmt.Sprintf("[GET /api/admin/{tid}/scopes/{scp}][%d] getScopeForbidden  %+v", 403, o.Payload)
}
func (o *GetScopeForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetScopeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScopeNotFound creates a GetScopeNotFound with default headers values
func NewGetScopeNotFound() *GetScopeNotFound {
	return &GetScopeNotFound{}
}

/* GetScopeNotFound describes a response with status code 404, with default header values.

HttpError
*/
type GetScopeNotFound struct {
	Payload *models.Error
}

func (o *GetScopeNotFound) Error() string {
	return fmt.Sprintf("[GET /api/admin/{tid}/scopes/{scp}][%d] getScopeNotFound  %+v", 404, o.Payload)
}
func (o *GetScopeNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetScopeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
