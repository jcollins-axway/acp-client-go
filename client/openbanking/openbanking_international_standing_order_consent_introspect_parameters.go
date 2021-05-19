// Code generated by go-swagger; DO NOT EDIT.

package openbanking

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

// NewOpenbankingInternationalStandingOrderConsentIntrospectParams creates a new OpenbankingInternationalStandingOrderConsentIntrospectParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewOpenbankingInternationalStandingOrderConsentIntrospectParams() *OpenbankingInternationalStandingOrderConsentIntrospectParams {
	return &OpenbankingInternationalStandingOrderConsentIntrospectParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewOpenbankingInternationalStandingOrderConsentIntrospectParamsWithTimeout creates a new OpenbankingInternationalStandingOrderConsentIntrospectParams object
// with the ability to set a timeout on a request.
func NewOpenbankingInternationalStandingOrderConsentIntrospectParamsWithTimeout(timeout time.Duration) *OpenbankingInternationalStandingOrderConsentIntrospectParams {
	return &OpenbankingInternationalStandingOrderConsentIntrospectParams{
		timeout: timeout,
	}
}

// NewOpenbankingInternationalStandingOrderConsentIntrospectParamsWithContext creates a new OpenbankingInternationalStandingOrderConsentIntrospectParams object
// with the ability to set a context for a request.
func NewOpenbankingInternationalStandingOrderConsentIntrospectParamsWithContext(ctx context.Context) *OpenbankingInternationalStandingOrderConsentIntrospectParams {
	return &OpenbankingInternationalStandingOrderConsentIntrospectParams{
		Context: ctx,
	}
}

// NewOpenbankingInternationalStandingOrderConsentIntrospectParamsWithHTTPClient creates a new OpenbankingInternationalStandingOrderConsentIntrospectParams object
// with the ability to set a custom HTTPClient for a request.
func NewOpenbankingInternationalStandingOrderConsentIntrospectParamsWithHTTPClient(client *http.Client) *OpenbankingInternationalStandingOrderConsentIntrospectParams {
	return &OpenbankingInternationalStandingOrderConsentIntrospectParams{
		HTTPClient: client,
	}
}

/* OpenbankingInternationalStandingOrderConsentIntrospectParams contains all the parameters to send to the API endpoint
   for the openbanking international standing order consent introspect operation.

   Typically these are written to a http.Request.
*/
type OpenbankingInternationalStandingOrderConsentIntrospectParams struct {

	/* Aid.

	   Authorization server id

	   Default: "default"
	*/
	Aid string

	/* Tid.

	   Tenant id

	   Default: "default"
	*/
	Tid string

	// Token.
	Token *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the openbanking international standing order consent introspect params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OpenbankingInternationalStandingOrderConsentIntrospectParams) WithDefaults() *OpenbankingInternationalStandingOrderConsentIntrospectParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the openbanking international standing order consent introspect params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OpenbankingInternationalStandingOrderConsentIntrospectParams) SetDefaults() {
	var (
		aidDefault = string("default")

		tidDefault = string("default")
	)

	val := OpenbankingInternationalStandingOrderConsentIntrospectParams{
		Aid: aidDefault,
		Tid: tidDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the openbanking international standing order consent introspect params
func (o *OpenbankingInternationalStandingOrderConsentIntrospectParams) WithTimeout(timeout time.Duration) *OpenbankingInternationalStandingOrderConsentIntrospectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the openbanking international standing order consent introspect params
func (o *OpenbankingInternationalStandingOrderConsentIntrospectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the openbanking international standing order consent introspect params
func (o *OpenbankingInternationalStandingOrderConsentIntrospectParams) WithContext(ctx context.Context) *OpenbankingInternationalStandingOrderConsentIntrospectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the openbanking international standing order consent introspect params
func (o *OpenbankingInternationalStandingOrderConsentIntrospectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the openbanking international standing order consent introspect params
func (o *OpenbankingInternationalStandingOrderConsentIntrospectParams) WithHTTPClient(client *http.Client) *OpenbankingInternationalStandingOrderConsentIntrospectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the openbanking international standing order consent introspect params
func (o *OpenbankingInternationalStandingOrderConsentIntrospectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAid adds the aid to the openbanking international standing order consent introspect params
func (o *OpenbankingInternationalStandingOrderConsentIntrospectParams) WithAid(aid string) *OpenbankingInternationalStandingOrderConsentIntrospectParams {
	o.SetAid(aid)
	return o
}

// SetAid adds the aid to the openbanking international standing order consent introspect params
func (o *OpenbankingInternationalStandingOrderConsentIntrospectParams) SetAid(aid string) {
	o.Aid = aid
}

// WithTid adds the tid to the openbanking international standing order consent introspect params
func (o *OpenbankingInternationalStandingOrderConsentIntrospectParams) WithTid(tid string) *OpenbankingInternationalStandingOrderConsentIntrospectParams {
	o.SetTid(tid)
	return o
}

// SetTid adds the tid to the openbanking international standing order consent introspect params
func (o *OpenbankingInternationalStandingOrderConsentIntrospectParams) SetTid(tid string) {
	o.Tid = tid
}

// WithToken adds the token to the openbanking international standing order consent introspect params
func (o *OpenbankingInternationalStandingOrderConsentIntrospectParams) WithToken(token *string) *OpenbankingInternationalStandingOrderConsentIntrospectParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the openbanking international standing order consent introspect params
func (o *OpenbankingInternationalStandingOrderConsentIntrospectParams) SetToken(token *string) {
	o.Token = token
}

// WriteToRequest writes these params to a swagger request
func (o *OpenbankingInternationalStandingOrderConsentIntrospectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param aid
	if err := r.SetPathParam("aid", o.Aid); err != nil {
		return err
	}

	// path param tid
	if err := r.SetPathParam("tid", o.Tid); err != nil {
		return err
	}

	if o.Token != nil {

		// form param token
		var frToken string
		if o.Token != nil {
			frToken = *o.Token
		}
		fToken := frToken
		if fToken != "" {
			if err := r.SetFormParam("token", fToken); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
