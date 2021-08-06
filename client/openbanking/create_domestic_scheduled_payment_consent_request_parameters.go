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

	"github.com/cloudentity/acp-client-go/models"
)

// NewCreateDomesticScheduledPaymentConsentRequestParams creates a new CreateDomesticScheduledPaymentConsentRequestParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateDomesticScheduledPaymentConsentRequestParams() *CreateDomesticScheduledPaymentConsentRequestParams {
	return &CreateDomesticScheduledPaymentConsentRequestParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateDomesticScheduledPaymentConsentRequestParamsWithTimeout creates a new CreateDomesticScheduledPaymentConsentRequestParams object
// with the ability to set a timeout on a request.
func NewCreateDomesticScheduledPaymentConsentRequestParamsWithTimeout(timeout time.Duration) *CreateDomesticScheduledPaymentConsentRequestParams {
	return &CreateDomesticScheduledPaymentConsentRequestParams{
		timeout: timeout,
	}
}

// NewCreateDomesticScheduledPaymentConsentRequestParamsWithContext creates a new CreateDomesticScheduledPaymentConsentRequestParams object
// with the ability to set a context for a request.
func NewCreateDomesticScheduledPaymentConsentRequestParamsWithContext(ctx context.Context) *CreateDomesticScheduledPaymentConsentRequestParams {
	return &CreateDomesticScheduledPaymentConsentRequestParams{
		Context: ctx,
	}
}

// NewCreateDomesticScheduledPaymentConsentRequestParamsWithHTTPClient creates a new CreateDomesticScheduledPaymentConsentRequestParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateDomesticScheduledPaymentConsentRequestParamsWithHTTPClient(client *http.Client) *CreateDomesticScheduledPaymentConsentRequestParams {
	return &CreateDomesticScheduledPaymentConsentRequestParams{
		HTTPClient: client,
	}
}

/* CreateDomesticScheduledPaymentConsentRequestParams contains all the parameters to send to the API endpoint
   for the create domestic scheduled payment consent request operation.

   Typically these are written to a http.Request.
*/
type CreateDomesticScheduledPaymentConsentRequestParams struct {

	/* Request.

	   Request
	*/
	Request *models.DomesticScheduledPaymentConsentRequest

	/* Aid.

	   Server ID

	   Default: "default"
	*/
	Aid string

	/* Tid.

	   Tenant ID

	   Default: "default"
	*/
	Tid string

	/* XCustomerUserAgent.

	     The header indicates the user-agent that the PSU is using.

	The TPP may populate this field with the user-agent indicated by the PSU.
	If the PSU is using a TPP mobile app, the TPP must ensure that the user-agent string
	is different from browser based user-agent strings.
	*/
	XCustomerUserAgent *string

	/* XFapiAuthDate.

	     The time when the PSU last logged in with the TPP.

	The value is supplied as a HTTP-date as in section 7.1.1.1 of [RFC7231]
	*/
	XFapiAuthDate *string

	/* XFapiCustomerIPAddress.

	   The PSU's IP address if the PSU is currently logged in with the TPP.
	*/
	XFapiCustomerIPAddress *string

	/* XFapiInteractionID.

	     An RFC4122 UID used as a correlation Id.

	If provided, the ASPSP must "play back" this value
	in the x-fapi-interaction-id response header.
	*/
	XFapiInteractionID *string

	/* XIdempotencyKey.

	     Every request will be processed only once per x-idempotency-key.
	The Idempotency Key will be valid for 24 hours
	*/
	XIdempotencyKey *string

	/* XJwsSignature.

	     Header containing a detached JWS signature of the body of the payload.

	Refer to resource specific documentation on when this header must be specified.
	*/
	XJwsSignature *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create domestic scheduled payment consent request params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateDomesticScheduledPaymentConsentRequestParams) WithDefaults() *CreateDomesticScheduledPaymentConsentRequestParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create domestic scheduled payment consent request params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateDomesticScheduledPaymentConsentRequestParams) SetDefaults() {
	var (
		aidDefault = string("default")

		tidDefault = string("default")
	)

	val := CreateDomesticScheduledPaymentConsentRequestParams{
		Aid: aidDefault,
		Tid: tidDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the create domestic scheduled payment consent request params
func (o *CreateDomesticScheduledPaymentConsentRequestParams) WithTimeout(timeout time.Duration) *CreateDomesticScheduledPaymentConsentRequestParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create domestic scheduled payment consent request params
func (o *CreateDomesticScheduledPaymentConsentRequestParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create domestic scheduled payment consent request params
func (o *CreateDomesticScheduledPaymentConsentRequestParams) WithContext(ctx context.Context) *CreateDomesticScheduledPaymentConsentRequestParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create domestic scheduled payment consent request params
func (o *CreateDomesticScheduledPaymentConsentRequestParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create domestic scheduled payment consent request params
func (o *CreateDomesticScheduledPaymentConsentRequestParams) WithHTTPClient(client *http.Client) *CreateDomesticScheduledPaymentConsentRequestParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create domestic scheduled payment consent request params
func (o *CreateDomesticScheduledPaymentConsentRequestParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the create domestic scheduled payment consent request params
func (o *CreateDomesticScheduledPaymentConsentRequestParams) WithRequest(request *models.DomesticScheduledPaymentConsentRequest) *CreateDomesticScheduledPaymentConsentRequestParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the create domestic scheduled payment consent request params
func (o *CreateDomesticScheduledPaymentConsentRequestParams) SetRequest(request *models.DomesticScheduledPaymentConsentRequest) {
	o.Request = request
}

// WithAid adds the aid to the create domestic scheduled payment consent request params
func (o *CreateDomesticScheduledPaymentConsentRequestParams) WithAid(aid string) *CreateDomesticScheduledPaymentConsentRequestParams {
	o.SetAid(aid)
	return o
}

// SetAid adds the aid to the create domestic scheduled payment consent request params
func (o *CreateDomesticScheduledPaymentConsentRequestParams) SetAid(aid string) {
	o.Aid = aid
}

// WithTid adds the tid to the create domestic scheduled payment consent request params
func (o *CreateDomesticScheduledPaymentConsentRequestParams) WithTid(tid string) *CreateDomesticScheduledPaymentConsentRequestParams {
	o.SetTid(tid)
	return o
}

// SetTid adds the tid to the create domestic scheduled payment consent request params
func (o *CreateDomesticScheduledPaymentConsentRequestParams) SetTid(tid string) {
	o.Tid = tid
}

// WithXCustomerUserAgent adds the xCustomerUserAgent to the create domestic scheduled payment consent request params
func (o *CreateDomesticScheduledPaymentConsentRequestParams) WithXCustomerUserAgent(xCustomerUserAgent *string) *CreateDomesticScheduledPaymentConsentRequestParams {
	o.SetXCustomerUserAgent(xCustomerUserAgent)
	return o
}

// SetXCustomerUserAgent adds the xCustomerUserAgent to the create domestic scheduled payment consent request params
func (o *CreateDomesticScheduledPaymentConsentRequestParams) SetXCustomerUserAgent(xCustomerUserAgent *string) {
	o.XCustomerUserAgent = xCustomerUserAgent
}

// WithXFapiAuthDate adds the xFapiAuthDate to the create domestic scheduled payment consent request params
func (o *CreateDomesticScheduledPaymentConsentRequestParams) WithXFapiAuthDate(xFapiAuthDate *string) *CreateDomesticScheduledPaymentConsentRequestParams {
	o.SetXFapiAuthDate(xFapiAuthDate)
	return o
}

// SetXFapiAuthDate adds the xFapiAuthDate to the create domestic scheduled payment consent request params
func (o *CreateDomesticScheduledPaymentConsentRequestParams) SetXFapiAuthDate(xFapiAuthDate *string) {
	o.XFapiAuthDate = xFapiAuthDate
}

// WithXFapiCustomerIPAddress adds the xFapiCustomerIPAddress to the create domestic scheduled payment consent request params
func (o *CreateDomesticScheduledPaymentConsentRequestParams) WithXFapiCustomerIPAddress(xFapiCustomerIPAddress *string) *CreateDomesticScheduledPaymentConsentRequestParams {
	o.SetXFapiCustomerIPAddress(xFapiCustomerIPAddress)
	return o
}

// SetXFapiCustomerIPAddress adds the xFapiCustomerIpAddress to the create domestic scheduled payment consent request params
func (o *CreateDomesticScheduledPaymentConsentRequestParams) SetXFapiCustomerIPAddress(xFapiCustomerIPAddress *string) {
	o.XFapiCustomerIPAddress = xFapiCustomerIPAddress
}

// WithXFapiInteractionID adds the xFapiInteractionID to the create domestic scheduled payment consent request params
func (o *CreateDomesticScheduledPaymentConsentRequestParams) WithXFapiInteractionID(xFapiInteractionID *string) *CreateDomesticScheduledPaymentConsentRequestParams {
	o.SetXFapiInteractionID(xFapiInteractionID)
	return o
}

// SetXFapiInteractionID adds the xFapiInteractionId to the create domestic scheduled payment consent request params
func (o *CreateDomesticScheduledPaymentConsentRequestParams) SetXFapiInteractionID(xFapiInteractionID *string) {
	o.XFapiInteractionID = xFapiInteractionID
}

// WithXIdempotencyKey adds the xIdempotencyKey to the create domestic scheduled payment consent request params
func (o *CreateDomesticScheduledPaymentConsentRequestParams) WithXIdempotencyKey(xIdempotencyKey *string) *CreateDomesticScheduledPaymentConsentRequestParams {
	o.SetXIdempotencyKey(xIdempotencyKey)
	return o
}

// SetXIdempotencyKey adds the xIdempotencyKey to the create domestic scheduled payment consent request params
func (o *CreateDomesticScheduledPaymentConsentRequestParams) SetXIdempotencyKey(xIdempotencyKey *string) {
	o.XIdempotencyKey = xIdempotencyKey
}

// WithXJwsSignature adds the xJwsSignature to the create domestic scheduled payment consent request params
func (o *CreateDomesticScheduledPaymentConsentRequestParams) WithXJwsSignature(xJwsSignature *string) *CreateDomesticScheduledPaymentConsentRequestParams {
	o.SetXJwsSignature(xJwsSignature)
	return o
}

// SetXJwsSignature adds the xJwsSignature to the create domestic scheduled payment consent request params
func (o *CreateDomesticScheduledPaymentConsentRequestParams) SetXJwsSignature(xJwsSignature *string) {
	o.XJwsSignature = xJwsSignature
}

// WriteToRequest writes these params to a swagger request
func (o *CreateDomesticScheduledPaymentConsentRequestParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Request != nil {
		if err := r.SetBodyParam(o.Request); err != nil {
			return err
		}
	}

	// path param aid
	if err := r.SetPathParam("aid", o.Aid); err != nil {
		return err
	}

	// path param tid
	if err := r.SetPathParam("tid", o.Tid); err != nil {
		return err
	}

	if o.XCustomerUserAgent != nil {

		// header param x-customer-user-agent
		if err := r.SetHeaderParam("x-customer-user-agent", *o.XCustomerUserAgent); err != nil {
			return err
		}
	}

	if o.XFapiAuthDate != nil {

		// header param x-fapi-auth-date
		if err := r.SetHeaderParam("x-fapi-auth-date", *o.XFapiAuthDate); err != nil {
			return err
		}
	}

	if o.XFapiCustomerIPAddress != nil {

		// header param x-fapi-customer-ip-address
		if err := r.SetHeaderParam("x-fapi-customer-ip-address", *o.XFapiCustomerIPAddress); err != nil {
			return err
		}
	}

	if o.XFapiInteractionID != nil {

		// header param x-fapi-interaction-id
		if err := r.SetHeaderParam("x-fapi-interaction-id", *o.XFapiInteractionID); err != nil {
			return err
		}
	}

	if o.XIdempotencyKey != nil {

		// header param x-idempotency-key
		if err := r.SetHeaderParam("x-idempotency-key", *o.XIdempotencyKey); err != nil {
			return err
		}
	}

	if o.XJwsSignature != nil {

		// header param x-jws-signature
		if err := r.SetHeaderParam("x-jws-signature", *o.XJwsSignature); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
