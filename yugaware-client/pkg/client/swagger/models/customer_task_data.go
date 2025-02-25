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

// CustomerTaskData Customer task data
//
// swagger:model CustomerTaskData
type CustomerTaskData struct {

	// Customer task completion time
	// Example: 1624295417405
	CompletionTime int64 `json:"completionTime,omitempty"`

	// Customer task creation time
	// Example: 1624295417405
	CreateTime int64 `json:"createTime,omitempty"`

	// Customer task UUID
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// Customer task percentage completed
	// Example: 100
	PercentComplete int32 `json:"percentComplete,omitempty"`

	// Customer task status
	// Example: Complete
	Status string `json:"status,omitempty"`

	// Customer task target
	// Example: Universe
	Target string `json:"target,omitempty"`

	// Customer task target UUID
	// Format: uuid
	TargetUUID strfmt.UUID `json:"targetUUID,omitempty"`

	// Customer task title
	// Example: Deleted Universe : test-universe
	Title string `json:"title,omitempty"`

	// Customer task type
	// Example: Delete
	Type string `json:"type,omitempty"`
}

// Validate validates this customer task data
func (m *CustomerTaskData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargetUUID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CustomerTaskData) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CustomerTaskData) validateTargetUUID(formats strfmt.Registry) error {
	if swag.IsZero(m.TargetUUID) { // not required
		return nil
	}

	if err := validate.FormatOf("targetUUID", "body", "uuid", m.TargetUUID.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this customer task data based on context it is used
func (m *CustomerTaskData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CustomerTaskData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustomerTaskData) UnmarshalBinary(b []byte) error {
	var res CustomerTaskData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
