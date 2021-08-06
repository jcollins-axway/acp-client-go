// Code generated by go-swagger; DO NOT EDIT.

package services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewImportSpecificationFromFileParams creates a new ImportSpecificationFromFileParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewImportSpecificationFromFileParams() *ImportSpecificationFromFileParams {
	return &ImportSpecificationFromFileParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewImportSpecificationFromFileParamsWithTimeout creates a new ImportSpecificationFromFileParams object
// with the ability to set a timeout on a request.
func NewImportSpecificationFromFileParamsWithTimeout(timeout time.Duration) *ImportSpecificationFromFileParams {
	return &ImportSpecificationFromFileParams{
		timeout: timeout,
	}
}

// NewImportSpecificationFromFileParamsWithContext creates a new ImportSpecificationFromFileParams object
// with the ability to set a context for a request.
func NewImportSpecificationFromFileParamsWithContext(ctx context.Context) *ImportSpecificationFromFileParams {
	return &ImportSpecificationFromFileParams{
		Context: ctx,
	}
}

// NewImportSpecificationFromFileParamsWithHTTPClient creates a new ImportSpecificationFromFileParams object
// with the ability to set a custom HTTPClient for a request.
func NewImportSpecificationFromFileParamsWithHTTPClient(client *http.Client) *ImportSpecificationFromFileParams {
	return &ImportSpecificationFromFileParams{
		HTTPClient: client,
	}
}

/* ImportSpecificationFromFileParams contains all the parameters to send to the API endpoint
   for the import specification from file operation.

   Typically these are written to a http.Request.
*/
type ImportSpecificationFromFileParams struct {

	// File.
	File runtime.NamedReadCloser

	// GatewayType.
	GatewayType *string

	// Sid.
	Sid string

	/* Tid.

	   Tenant id

	   Default: "default"
	*/
	Tid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the import specification from file params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ImportSpecificationFromFileParams) WithDefaults() *ImportSpecificationFromFileParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the import specification from file params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ImportSpecificationFromFileParams) SetDefaults() {
	var (
		tidDefault = string("default")
	)

	val := ImportSpecificationFromFileParams{
		Tid: tidDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the import specification from file params
func (o *ImportSpecificationFromFileParams) WithTimeout(timeout time.Duration) *ImportSpecificationFromFileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the import specification from file params
func (o *ImportSpecificationFromFileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the import specification from file params
func (o *ImportSpecificationFromFileParams) WithContext(ctx context.Context) *ImportSpecificationFromFileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the import specification from file params
func (o *ImportSpecificationFromFileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the import specification from file params
func (o *ImportSpecificationFromFileParams) WithHTTPClient(client *http.Client) *ImportSpecificationFromFileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the import specification from file params
func (o *ImportSpecificationFromFileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFile adds the file to the import specification from file params
func (o *ImportSpecificationFromFileParams) WithFile(file runtime.NamedReadCloser) *ImportSpecificationFromFileParams {
	o.SetFile(file)
	return o
}

// SetFile adds the file to the import specification from file params
func (o *ImportSpecificationFromFileParams) SetFile(file runtime.NamedReadCloser) {
	o.File = file
}

// WithGatewayType adds the gatewayType to the import specification from file params
func (o *ImportSpecificationFromFileParams) WithGatewayType(gatewayType *string) *ImportSpecificationFromFileParams {
	o.SetGatewayType(gatewayType)
	return o
}

// SetGatewayType adds the gatewayType to the import specification from file params
func (o *ImportSpecificationFromFileParams) SetGatewayType(gatewayType *string) {
	o.GatewayType = gatewayType
}

// WithSid adds the sid to the import specification from file params
func (o *ImportSpecificationFromFileParams) WithSid(sid string) *ImportSpecificationFromFileParams {
	o.SetSid(sid)
	return o
}

// SetSid adds the sid to the import specification from file params
func (o *ImportSpecificationFromFileParams) SetSid(sid string) {
	o.Sid = sid
}

// WithTid adds the tid to the import specification from file params
func (o *ImportSpecificationFromFileParams) WithTid(tid string) *ImportSpecificationFromFileParams {
	o.SetTid(tid)
	return o
}

// SetTid adds the tid to the import specification from file params
func (o *ImportSpecificationFromFileParams) SetTid(tid string) {
	o.Tid = tid
}

// WriteToRequest writes these params to a swagger request
func (o *ImportSpecificationFromFileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.File != nil {

		if o.File != nil {
			// form file param file
			if err := r.SetFileParam("file", o.File); err != nil {
				return err
			}
		}
	}

	if o.GatewayType != nil {

		// query param gateway_type
		var qrGatewayType string

		if o.GatewayType != nil {
			qrGatewayType = *o.GatewayType
		}
		qGatewayType := qrGatewayType
		if qGatewayType != "" {

			if err := r.SetQueryParam("gateway_type", qGatewayType); err != nil {
				return err
			}
		}
	}

	// path param sid
	if err := r.SetPathParam("sid", o.Sid); err != nil {
		return err
	}

	// path param tid
	if err := r.SetPathParam("tid", o.Tid); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
