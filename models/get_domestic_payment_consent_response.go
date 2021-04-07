// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetDomesticPaymentConsentResponse get domestic payment consent response
//
// swagger:model GetDomesticPaymentConsentResponse
type GetDomesticPaymentConsentResponse struct {

	// todo added here for simplification before generalization effort
	AccountIDs []string `json:"AccountIDs"`

	// authorisation
	Authorisation *DomesticPaymentConsentAuthorisation `json:"Authorisation,omitempty"`

	// Unique identification as assigned to identify the domestic payment resource.
	ConsentID string `json:"consent_id,omitempty"`

	// Date and time at which the resource was created.
	// Format: date-time
	CreationDateTime strfmt.DateTime `json:"CreationDateTime,omitempty"`

	// delivery address
	DeliveryAddress *RiskDeliveryAddress `json:"DeliveryAddress,omitempty"`

	// initiation
	Initiation *DomesticPaymentConsentDataInitiation `json:"Initiation,omitempty"`

	// Category code conform to ISO 18245, related to the type of services or goods the merchant provides for the transaction.
	// Max Length: 4
	// Min Length: 3
	MerchantCategoryCode string `json:"MerchantCategoryCode,omitempty"`

	// The unique customer identifier of the PSU with the merchant.
	// Max Length: 70
	// Min Length: 1
	MerchantCustomerIdentification string `json:"MerchantCustomerIdentification,omitempty"`

	// Specifies the payment context
	// Enum: [[BillPayment EcommerceGoods EcommerceServices Other PartyToParty]]
	PaymentContextCode string `json:"PaymentContextCode,omitempty"`

	// Specifies to share the refund account details with PISP
	// Enum: [[No Yes]]
	ReadRefundAccount string `json:"ReadRefundAccount,omitempty"`

	// requested scopes
	RequestedScopes []*RequestedScope `json:"requested_scopes"`

	// s c a support data
	SCASupportData *DomesticPaymentConsentSCASupportData `json:"SCASupportData,omitempty"`

	// Specifies the status of consent resource in code form.
	Status string `json:"Status,omitempty"`

	// Date and time at which the resource status was updated.
	// Format: date-time
	StatusUpdateDateTime strfmt.DateTime `json:"StatusUpdateDateTime,omitempty"`

	// subject
	Subject string `json:"subject,omitempty"`

	// authentication context
	AuthenticationContext AuthenticationContext `json:"authentication_context,omitempty"`

	// client
	Client *ClientInfo `json:"client,omitempty"`
}

// Validate validates this get domestic payment consent response
func (m *GetDomesticPaymentConsentResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthorisation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreationDateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeliveryAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInitiation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMerchantCategoryCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMerchantCustomerIdentification(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePaymentContextCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReadRefundAccount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequestedScopes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSCASupportData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatusUpdateDateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuthenticationContext(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClient(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetDomesticPaymentConsentResponse) validateAuthorisation(formats strfmt.Registry) error {
	if swag.IsZero(m.Authorisation) { // not required
		return nil
	}

	if m.Authorisation != nil {
		if err := m.Authorisation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Authorisation")
			}
			return err
		}
	}

	return nil
}

func (m *GetDomesticPaymentConsentResponse) validateCreationDateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.CreationDateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("CreationDateTime", "body", "date-time", m.CreationDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *GetDomesticPaymentConsentResponse) validateDeliveryAddress(formats strfmt.Registry) error {
	if swag.IsZero(m.DeliveryAddress) { // not required
		return nil
	}

	if m.DeliveryAddress != nil {
		if err := m.DeliveryAddress.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("DeliveryAddress")
			}
			return err
		}
	}

	return nil
}

func (m *GetDomesticPaymentConsentResponse) validateInitiation(formats strfmt.Registry) error {
	if swag.IsZero(m.Initiation) { // not required
		return nil
	}

	if m.Initiation != nil {
		if err := m.Initiation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Initiation")
			}
			return err
		}
	}

	return nil
}

func (m *GetDomesticPaymentConsentResponse) validateMerchantCategoryCode(formats strfmt.Registry) error {
	if swag.IsZero(m.MerchantCategoryCode) { // not required
		return nil
	}

	if err := validate.MinLength("MerchantCategoryCode", "body", m.MerchantCategoryCode, 3); err != nil {
		return err
	}

	if err := validate.MaxLength("MerchantCategoryCode", "body", m.MerchantCategoryCode, 4); err != nil {
		return err
	}

	return nil
}

func (m *GetDomesticPaymentConsentResponse) validateMerchantCustomerIdentification(formats strfmt.Registry) error {
	if swag.IsZero(m.MerchantCustomerIdentification) { // not required
		return nil
	}

	if err := validate.MinLength("MerchantCustomerIdentification", "body", m.MerchantCustomerIdentification, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("MerchantCustomerIdentification", "body", m.MerchantCustomerIdentification, 70); err != nil {
		return err
	}

	return nil
}

var getDomesticPaymentConsentResponseTypePaymentContextCodePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["[BillPayment EcommerceGoods EcommerceServices Other PartyToParty]"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getDomesticPaymentConsentResponseTypePaymentContextCodePropEnum = append(getDomesticPaymentConsentResponseTypePaymentContextCodePropEnum, v)
	}
}

const (

	// GetDomesticPaymentConsentResponsePaymentContextCodeBillPaymentEcommerceGoodsEcommerceServicesOtherPartyToParty captures enum value "[BillPayment EcommerceGoods EcommerceServices Other PartyToParty]"
	GetDomesticPaymentConsentResponsePaymentContextCodeBillPaymentEcommerceGoodsEcommerceServicesOtherPartyToParty string = "[BillPayment EcommerceGoods EcommerceServices Other PartyToParty]"
)

// prop value enum
func (m *GetDomesticPaymentConsentResponse) validatePaymentContextCodeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, getDomesticPaymentConsentResponseTypePaymentContextCodePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *GetDomesticPaymentConsentResponse) validatePaymentContextCode(formats strfmt.Registry) error {
	if swag.IsZero(m.PaymentContextCode) { // not required
		return nil
	}

	// value enum
	if err := m.validatePaymentContextCodeEnum("PaymentContextCode", "body", m.PaymentContextCode); err != nil {
		return err
	}

	return nil
}

var getDomesticPaymentConsentResponseTypeReadRefundAccountPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["[No Yes]"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getDomesticPaymentConsentResponseTypeReadRefundAccountPropEnum = append(getDomesticPaymentConsentResponseTypeReadRefundAccountPropEnum, v)
	}
}

const (

	// GetDomesticPaymentConsentResponseReadRefundAccountNoYes captures enum value "[No Yes]"
	GetDomesticPaymentConsentResponseReadRefundAccountNoYes string = "[No Yes]"
)

// prop value enum
func (m *GetDomesticPaymentConsentResponse) validateReadRefundAccountEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, getDomesticPaymentConsentResponseTypeReadRefundAccountPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *GetDomesticPaymentConsentResponse) validateReadRefundAccount(formats strfmt.Registry) error {
	if swag.IsZero(m.ReadRefundAccount) { // not required
		return nil
	}

	// value enum
	if err := m.validateReadRefundAccountEnum("ReadRefundAccount", "body", m.ReadRefundAccount); err != nil {
		return err
	}

	return nil
}

func (m *GetDomesticPaymentConsentResponse) validateRequestedScopes(formats strfmt.Registry) error {
	if swag.IsZero(m.RequestedScopes) { // not required
		return nil
	}

	for i := 0; i < len(m.RequestedScopes); i++ {
		if swag.IsZero(m.RequestedScopes[i]) { // not required
			continue
		}

		if m.RequestedScopes[i] != nil {
			if err := m.RequestedScopes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("requested_scopes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *GetDomesticPaymentConsentResponse) validateSCASupportData(formats strfmt.Registry) error {
	if swag.IsZero(m.SCASupportData) { // not required
		return nil
	}

	if m.SCASupportData != nil {
		if err := m.SCASupportData.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("SCASupportData")
			}
			return err
		}
	}

	return nil
}

func (m *GetDomesticPaymentConsentResponse) validateStatusUpdateDateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.StatusUpdateDateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("StatusUpdateDateTime", "body", "date-time", m.StatusUpdateDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *GetDomesticPaymentConsentResponse) validateAuthenticationContext(formats strfmt.Registry) error {
	if swag.IsZero(m.AuthenticationContext) { // not required
		return nil
	}

	if m.AuthenticationContext != nil {
		if err := m.AuthenticationContext.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("authentication_context")
			}
			return err
		}
	}

	return nil
}

func (m *GetDomesticPaymentConsentResponse) validateClient(formats strfmt.Registry) error {
	if swag.IsZero(m.Client) { // not required
		return nil
	}

	if m.Client != nil {
		if err := m.Client.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("client")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get domestic payment consent response based on the context it is used
func (m *GetDomesticPaymentConsentResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAuthorisation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDeliveryAddress(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInitiation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRequestedScopes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSCASupportData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAuthenticationContext(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateClient(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetDomesticPaymentConsentResponse) contextValidateAuthorisation(ctx context.Context, formats strfmt.Registry) error {

	if m.Authorisation != nil {
		if err := m.Authorisation.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Authorisation")
			}
			return err
		}
	}

	return nil
}

func (m *GetDomesticPaymentConsentResponse) contextValidateDeliveryAddress(ctx context.Context, formats strfmt.Registry) error {

	if m.DeliveryAddress != nil {
		if err := m.DeliveryAddress.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("DeliveryAddress")
			}
			return err
		}
	}

	return nil
}

func (m *GetDomesticPaymentConsentResponse) contextValidateInitiation(ctx context.Context, formats strfmt.Registry) error {

	if m.Initiation != nil {
		if err := m.Initiation.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Initiation")
			}
			return err
		}
	}

	return nil
}

func (m *GetDomesticPaymentConsentResponse) contextValidateRequestedScopes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.RequestedScopes); i++ {

		if m.RequestedScopes[i] != nil {
			if err := m.RequestedScopes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("requested_scopes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *GetDomesticPaymentConsentResponse) contextValidateSCASupportData(ctx context.Context, formats strfmt.Registry) error {

	if m.SCASupportData != nil {
		if err := m.SCASupportData.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("SCASupportData")
			}
			return err
		}
	}

	return nil
}

func (m *GetDomesticPaymentConsentResponse) contextValidateAuthenticationContext(ctx context.Context, formats strfmt.Registry) error {

	if err := m.AuthenticationContext.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("authentication_context")
		}
		return err
	}

	return nil
}

func (m *GetDomesticPaymentConsentResponse) contextValidateClient(ctx context.Context, formats strfmt.Registry) error {

	if m.Client != nil {
		if err := m.Client.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("client")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetDomesticPaymentConsentResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetDomesticPaymentConsentResponse) UnmarshalBinary(b []byte) error {
	var res GetDomesticPaymentConsentResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
