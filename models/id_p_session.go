// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// IDPSession ID p session
//
// swagger:model IDPSession
type IDPSession struct {

	// ID
	ID IDPSessionID `json:"ID,omitempty"`

	// idp ID
	IDPID string `json:"IDPID,omitempty"`

	// authentication context class reference
	Acr string `json:"acr,omitempty"`

	// authentication methods references
	Amr []string `json:"amr"`

	// time when user authenticated
	// Format: date-time
	AuthTime strfmt.DateTime `json:"auth_time,omitempty"`

	// authentication context
	AuthenticationContext AuthenticationContext `json:"authentication_context,omitempty"`

	// time when the session was issued
	// Format: date-time
	IssueTime strfmt.DateTime `json:"issue_time,omitempty"`

	// max age
	MaxAge Duration `json:"max_age,omitempty"`

	// user identifier
	// Example: user
	Subject string `json:"subject,omitempty"`
}

// Validate validates this ID p session
func (m *IDPSession) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuthTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuthenticationContext(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIssueTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaxAge(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IDPSession) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := m.ID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("ID")
		}
		return err
	}

	return nil
}

func (m *IDPSession) validateAuthTime(formats strfmt.Registry) error {
	if swag.IsZero(m.AuthTime) { // not required
		return nil
	}

	if err := validate.FormatOf("auth_time", "body", "date-time", m.AuthTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *IDPSession) validateAuthenticationContext(formats strfmt.Registry) error {
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

func (m *IDPSession) validateIssueTime(formats strfmt.Registry) error {
	if swag.IsZero(m.IssueTime) { // not required
		return nil
	}

	if err := validate.FormatOf("issue_time", "body", "date-time", m.IssueTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *IDPSession) validateMaxAge(formats strfmt.Registry) error {
	if swag.IsZero(m.MaxAge) { // not required
		return nil
	}

	if err := m.MaxAge.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("max_age")
		}
		return err
	}

	return nil
}

// ContextValidate validate this ID p session based on the context it is used
func (m *IDPSession) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAuthenticationContext(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMaxAge(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IDPSession) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := m.ID.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("ID")
		}
		return err
	}

	return nil
}

func (m *IDPSession) contextValidateAuthenticationContext(ctx context.Context, formats strfmt.Registry) error {

	if err := m.AuthenticationContext.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("authentication_context")
		}
		return err
	}

	return nil
}

func (m *IDPSession) contextValidateMaxAge(ctx context.Context, formats strfmt.Registry) error {

	if err := m.MaxAge.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("max_age")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IDPSession) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IDPSession) UnmarshalBinary(b []byte) error {
	var res IDPSession
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
