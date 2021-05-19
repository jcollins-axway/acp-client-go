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

// RecurringJob recurring job
//
// swagger:model RecurringJob
type RecurringJob struct {

	// cron expression
	Cron string `json:"cron,omitempty"`

	// id
	ID JobID `json:"id,omitempty"`

	// is paused
	Paused bool `json:"paused,omitempty"`

	// payload
	Payload interface{} `json:"payload,omitempty"`

	// queue
	Queue QueueName `json:"queue,omitempty"`

	// next execution time
	// Format: date-time
	ScheduledAt strfmt.DateTime `json:"scheduled_at,omitempty"`

	// job starting from
	// Format: date-time
	StartingFrom strfmt.DateTime `json:"starting_from,omitempty"`

	// tenant id
	TenantID string `json:"tenant_id,omitempty"`
}

// Validate validates this recurring job
func (m *RecurringJob) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQueue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScheduledAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartingFrom(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecurringJob) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := m.ID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("id")
		}
		return err
	}

	return nil
}

func (m *RecurringJob) validateQueue(formats strfmt.Registry) error {
	if swag.IsZero(m.Queue) { // not required
		return nil
	}

	if err := m.Queue.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("queue")
		}
		return err
	}

	return nil
}

func (m *RecurringJob) validateScheduledAt(formats strfmt.Registry) error {
	if swag.IsZero(m.ScheduledAt) { // not required
		return nil
	}

	if err := validate.FormatOf("scheduled_at", "body", "date-time", m.ScheduledAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RecurringJob) validateStartingFrom(formats strfmt.Registry) error {
	if swag.IsZero(m.StartingFrom) { // not required
		return nil
	}

	if err := validate.FormatOf("starting_from", "body", "date-time", m.StartingFrom.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this recurring job based on the context it is used
func (m *RecurringJob) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateQueue(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecurringJob) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := m.ID.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("id")
		}
		return err
	}

	return nil
}

func (m *RecurringJob) contextValidateQueue(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Queue.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("queue")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RecurringJob) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecurringJob) UnmarshalBinary(b []byte) error {
	var res RecurringJob
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
