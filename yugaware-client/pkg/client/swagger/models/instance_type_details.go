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

// InstanceTypeDetails instance type details
//
// swagger:model InstanceTypeDetails
type InstanceTypeDetails struct {

	// tenancy
	// Required: true
	// Enum: [Shared Dedicated Host]
	Tenancy *string `json:"tenancy"`

	// volume details list
	// Required: true
	VolumeDetailsList []*VolumeDetails `json:"volumeDetailsList"`
}

// Validate validates this instance type details
func (m *InstanceTypeDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTenancy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVolumeDetailsList(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var instanceTypeDetailsTypeTenancyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Shared","Dedicated","Host"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		instanceTypeDetailsTypeTenancyPropEnum = append(instanceTypeDetailsTypeTenancyPropEnum, v)
	}
}

const (

	// InstanceTypeDetailsTenancyShared captures enum value "Shared"
	InstanceTypeDetailsTenancyShared string = "Shared"

	// InstanceTypeDetailsTenancyDedicated captures enum value "Dedicated"
	InstanceTypeDetailsTenancyDedicated string = "Dedicated"

	// InstanceTypeDetailsTenancyHost captures enum value "Host"
	InstanceTypeDetailsTenancyHost string = "Host"
)

// prop value enum
func (m *InstanceTypeDetails) validateTenancyEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, instanceTypeDetailsTypeTenancyPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *InstanceTypeDetails) validateTenancy(formats strfmt.Registry) error {

	if err := validate.Required("tenancy", "body", m.Tenancy); err != nil {
		return err
	}

	// value enum
	if err := m.validateTenancyEnum("tenancy", "body", *m.Tenancy); err != nil {
		return err
	}

	return nil
}

func (m *InstanceTypeDetails) validateVolumeDetailsList(formats strfmt.Registry) error {

	if err := validate.Required("volumeDetailsList", "body", m.VolumeDetailsList); err != nil {
		return err
	}

	for i := 0; i < len(m.VolumeDetailsList); i++ {
		if swag.IsZero(m.VolumeDetailsList[i]) { // not required
			continue
		}

		if m.VolumeDetailsList[i] != nil {
			if err := m.VolumeDetailsList[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("volumeDetailsList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this instance type details based on the context it is used
func (m *InstanceTypeDetails) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateVolumeDetailsList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InstanceTypeDetails) contextValidateVolumeDetailsList(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.VolumeDetailsList); i++ {

		if m.VolumeDetailsList[i] != nil {
			if err := m.VolumeDetailsList[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("volumeDetailsList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *InstanceTypeDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InstanceTypeDetails) UnmarshalBinary(b []byte) error {
	var res InstanceTypeDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
