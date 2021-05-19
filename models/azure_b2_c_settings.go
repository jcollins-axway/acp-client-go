// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AzureB2CSettings Azure AD B2C authentication settings.
//
// Provide OAuth client details here.
//
// swagger:model AzureB2CSettings
type AzureB2CSettings struct {

	// OAuth client identifier
	// Example: client
	ClientID string `json:"client_id,omitempty"`

	// a flag to fetch user groups
	// Example: true
	FetchGroups bool `json:"fetch_groups,omitempty"`

	// flag to fetch additional user data from graph endpoint
	GetUser bool `json:"get_user,omitempty"`

	// list of user attributes to be fetched from graph
	GraphUserAttributes []string `json:"graph_user_attributes"`

	// user groups format: id or name
	// Example: id
	GroupNameFormat string `json:"group_name_format,omitempty"`

	// should only security groups be fetched
	// Example: true
	OnlySecurityGroups bool `json:"only_security_groups,omitempty"`

	// The user flow to be run.
	// Specify the name of a user flow you've created in your Azure AD B2C tenant.
	// Example: b2c_1_sign_in
	Policy string `json:"policy,omitempty"`

	// OAuth scopes which client will be requesting
	// Example: ["email","profile","openid"]
	Scopes []string `json:"scopes"`

	// azure tenant id
	// Example: 123-312-123
	Tenant string `json:"tenant,omitempty"`
}

// Validate validates this azure b2 c settings
func (m *AzureB2CSettings) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this azure b2 c settings based on context it is used
func (m *AzureB2CSettings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AzureB2CSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AzureB2CSettings) UnmarshalBinary(b []byte) error {
	var res AzureB2CSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
