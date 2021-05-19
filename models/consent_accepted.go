// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ConsentAccepted consent accepted
//
// swagger:model ConsentAccepted
type ConsentAccepted struct {

	// url where user should be redirected
	RedirectTo string `json:"redirect_to,omitempty"`
}

// Validate validates this consent accepted
func (m *ConsentAccepted) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this consent accepted based on context it is used
func (m *ConsentAccepted) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ConsentAccepted) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConsentAccepted) UnmarshalBinary(b []byte) error {
	var res ConsentAccepted
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
