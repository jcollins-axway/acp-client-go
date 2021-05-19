// Code generated by go-swagger; DO NOT EDIT.

package oauth2

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudentity/acp-client-go/models"
)

// JwksReader is a Reader for the Jwks structure.
type JwksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *JwksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewJwksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewJwksNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewJwksOK creates a JwksOK with default headers values
func NewJwksOK() *JwksOK {
	return &JwksOK{}
}

/* JwksOK describes a response with status code 200, with default header values.

ClientJWKs
*/
type JwksOK struct {
	Payload *models.ClientJWKs
}

func (o *JwksOK) Error() string {
	return fmt.Sprintf("[GET /{tid}/{aid}/.well-known/jwks.json][%d] jwksOK  %+v", 200, o.Payload)
}
func (o *JwksOK) GetPayload() *models.ClientJWKs {
	return o.Payload
}

func (o *JwksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClientJWKs)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewJwksNotFound creates a JwksNotFound with default headers values
func NewJwksNotFound() *JwksNotFound {
	return &JwksNotFound{}
}

/* JwksNotFound describes a response with status code 404, with default header values.

genericError
*/
type JwksNotFound struct {
	Payload *models.GenericError
}

func (o *JwksNotFound) Error() string {
	return fmt.Sprintf("[GET /{tid}/{aid}/.well-known/jwks.json][%d] jwksNotFound  %+v", 404, o.Payload)
}
func (o *JwksNotFound) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *JwksNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
