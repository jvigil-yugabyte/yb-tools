// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// XClusterReplicationFormData x cluster replication form data
//
// swagger:model XClusterReplicationFormData
type XClusterReplicationFormData struct {

	// bootstrap ids
	// Required: true
	BootstrapIds []string `json:"bootstrapIds"`

	// Amazon Resource Name (ARN) of the CMK
	CmkArn string `json:"cmkArn,omitempty"`

	// Communication ports
	CommunicationPorts *CommunicationPorts `json:"communicationPorts,omitempty"`

	// Device information
	DeviceInfo *DeviceInfo `json:"deviceInfo,omitempty"`

	// Encryption at rest configation
	EncryptionAtRestConfig *EncryptionAtRestConfig `json:"encryptionAtRestConfig,omitempty"`

	// Error message
	ErrorString string `json:"errorString,omitempty"`

	// Expected universe version
	ExpectedUniverseVersion int32 `json:"expectedUniverseVersion,omitempty"`

	// Extra dependencies
	ExtraDependencies *ExtraDependencies `json:"extraDependencies,omitempty"`

	// Whether this task has been tried before
	FirstTry bool `json:"firstTry,omitempty"`

	// Node details
	// Unique: true
	NodeDetailsSet []*NodeDetails `json:"nodeDetailsSet"`

	// Node exporter user
	NodeExporterUser string `json:"nodeExporterUser,omitempty"`

	// The source universe's sync replication relationships
	// Read Only: true
	SourceAsyncReplicationRelationships []*AsyncReplicationConfig `json:"sourceAsyncReplicationRelationships"`

	// source master addresses to change
	// Required: true
	SourceMasterAddressesToChange []string `json:"sourceMasterAddressesToChange"`

	// source table ids
	// Required: true
	SourceTableIds []string `json:"sourceTableIds"`

	// source table ids to add
	// Required: true
	SourceTableIdsToAdd []string `json:"sourceTableIdsToAdd"`

	// source table ids to remove
	// Required: true
	SourceTableIdsToRemove []string `json:"sourceTableIdsToRemove"`

	// source universe UUID
	// Required: true
	// Format: uuid
	SourceUniverseUUID *strfmt.UUID `json:"sourceUniverseUUID"`

	// The target universe's async replication relationships
	// Read Only: true
	TargetAsyncReplicationRelationships []*AsyncReplicationConfig `json:"targetAsyncReplicationRelationships"`

	// Associated universe UUID
	// Format: uuid
	UniverseUUID strfmt.UUID `json:"universeUUID,omitempty"`

	// Previous software version
	YbPrevSoftwareVersion string `json:"ybPrevSoftwareVersion,omitempty"`
}

// Validate validates this x cluster replication form data
func (m *XClusterReplicationFormData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBootstrapIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCommunicationPorts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeviceInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEncryptionAtRestConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExtraDependencies(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNodeDetailsSet(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceAsyncReplicationRelationships(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceMasterAddressesToChange(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceTableIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceTableIdsToAdd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceTableIdsToRemove(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceUniverseUUID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargetAsyncReplicationRelationships(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUniverseUUID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *XClusterReplicationFormData) validateBootstrapIds(formats strfmt.Registry) error {

	if err := validate.Required("bootstrapIds", "body", m.BootstrapIds); err != nil {
		return err
	}

	return nil
}

func (m *XClusterReplicationFormData) validateCommunicationPorts(formats strfmt.Registry) error {
	if swag.IsZero(m.CommunicationPorts) { // not required
		return nil
	}

	if m.CommunicationPorts != nil {
		if err := m.CommunicationPorts.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("communicationPorts")
			}
			return err
		}
	}

	return nil
}

func (m *XClusterReplicationFormData) validateDeviceInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.DeviceInfo) { // not required
		return nil
	}

	if m.DeviceInfo != nil {
		if err := m.DeviceInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deviceInfo")
			}
			return err
		}
	}

	return nil
}

func (m *XClusterReplicationFormData) validateEncryptionAtRestConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.EncryptionAtRestConfig) { // not required
		return nil
	}

	if m.EncryptionAtRestConfig != nil {
		if err := m.EncryptionAtRestConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("encryptionAtRestConfig")
			}
			return err
		}
	}

	return nil
}

func (m *XClusterReplicationFormData) validateExtraDependencies(formats strfmt.Registry) error {
	if swag.IsZero(m.ExtraDependencies) { // not required
		return nil
	}

	if m.ExtraDependencies != nil {
		if err := m.ExtraDependencies.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("extraDependencies")
			}
			return err
		}
	}

	return nil
}

func (m *XClusterReplicationFormData) validateNodeDetailsSet(formats strfmt.Registry) error {
	if swag.IsZero(m.NodeDetailsSet) { // not required
		return nil
	}

	if err := validate.UniqueItems("nodeDetailsSet", "body", m.NodeDetailsSet); err != nil {
		return err
	}

	for i := 0; i < len(m.NodeDetailsSet); i++ {
		if swag.IsZero(m.NodeDetailsSet[i]) { // not required
			continue
		}

		if m.NodeDetailsSet[i] != nil {
			if err := m.NodeDetailsSet[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("nodeDetailsSet" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *XClusterReplicationFormData) validateSourceAsyncReplicationRelationships(formats strfmt.Registry) error {
	if swag.IsZero(m.SourceAsyncReplicationRelationships) { // not required
		return nil
	}

	for i := 0; i < len(m.SourceAsyncReplicationRelationships); i++ {
		if swag.IsZero(m.SourceAsyncReplicationRelationships[i]) { // not required
			continue
		}

		if m.SourceAsyncReplicationRelationships[i] != nil {
			if err := m.SourceAsyncReplicationRelationships[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sourceAsyncReplicationRelationships" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *XClusterReplicationFormData) validateSourceMasterAddressesToChange(formats strfmt.Registry) error {

	if err := validate.Required("sourceMasterAddressesToChange", "body", m.SourceMasterAddressesToChange); err != nil {
		return err
	}

	return nil
}

func (m *XClusterReplicationFormData) validateSourceTableIds(formats strfmt.Registry) error {

	if err := validate.Required("sourceTableIds", "body", m.SourceTableIds); err != nil {
		return err
	}

	return nil
}

func (m *XClusterReplicationFormData) validateSourceTableIdsToAdd(formats strfmt.Registry) error {

	if err := validate.Required("sourceTableIdsToAdd", "body", m.SourceTableIdsToAdd); err != nil {
		return err
	}

	return nil
}

func (m *XClusterReplicationFormData) validateSourceTableIdsToRemove(formats strfmt.Registry) error {

	if err := validate.Required("sourceTableIdsToRemove", "body", m.SourceTableIdsToRemove); err != nil {
		return err
	}

	return nil
}

func (m *XClusterReplicationFormData) validateSourceUniverseUUID(formats strfmt.Registry) error {

	if err := validate.Required("sourceUniverseUUID", "body", m.SourceUniverseUUID); err != nil {
		return err
	}

	if err := validate.FormatOf("sourceUniverseUUID", "body", "uuid", m.SourceUniverseUUID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *XClusterReplicationFormData) validateTargetAsyncReplicationRelationships(formats strfmt.Registry) error {
	if swag.IsZero(m.TargetAsyncReplicationRelationships) { // not required
		return nil
	}

	for i := 0; i < len(m.TargetAsyncReplicationRelationships); i++ {
		if swag.IsZero(m.TargetAsyncReplicationRelationships[i]) { // not required
			continue
		}

		if m.TargetAsyncReplicationRelationships[i] != nil {
			if err := m.TargetAsyncReplicationRelationships[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("targetAsyncReplicationRelationships" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *XClusterReplicationFormData) validateUniverseUUID(formats strfmt.Registry) error {
	if swag.IsZero(m.UniverseUUID) { // not required
		return nil
	}

	if err := validate.FormatOf("universeUUID", "body", "uuid", m.UniverseUUID.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this x cluster replication form data based on the context it is used
func (m *XClusterReplicationFormData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCommunicationPorts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDeviceInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEncryptionAtRestConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExtraDependencies(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNodeDetailsSet(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSourceAsyncReplicationRelationships(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTargetAsyncReplicationRelationships(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *XClusterReplicationFormData) contextValidateCommunicationPorts(ctx context.Context, formats strfmt.Registry) error {

	if m.CommunicationPorts != nil {
		if err := m.CommunicationPorts.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("communicationPorts")
			}
			return err
		}
	}

	return nil
}

func (m *XClusterReplicationFormData) contextValidateDeviceInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.DeviceInfo != nil {
		if err := m.DeviceInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deviceInfo")
			}
			return err
		}
	}

	return nil
}

func (m *XClusterReplicationFormData) contextValidateEncryptionAtRestConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.EncryptionAtRestConfig != nil {
		if err := m.EncryptionAtRestConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("encryptionAtRestConfig")
			}
			return err
		}
	}

	return nil
}

func (m *XClusterReplicationFormData) contextValidateExtraDependencies(ctx context.Context, formats strfmt.Registry) error {

	if m.ExtraDependencies != nil {
		if err := m.ExtraDependencies.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("extraDependencies")
			}
			return err
		}
	}

	return nil
}

func (m *XClusterReplicationFormData) contextValidateNodeDetailsSet(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.NodeDetailsSet); i++ {

		if m.NodeDetailsSet[i] != nil {
			if err := m.NodeDetailsSet[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("nodeDetailsSet" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *XClusterReplicationFormData) contextValidateSourceAsyncReplicationRelationships(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "sourceAsyncReplicationRelationships", "body", []*AsyncReplicationConfig(m.SourceAsyncReplicationRelationships)); err != nil {
		return err
	}

	for i := 0; i < len(m.SourceAsyncReplicationRelationships); i++ {

		if m.SourceAsyncReplicationRelationships[i] != nil {
			if err := m.SourceAsyncReplicationRelationships[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sourceAsyncReplicationRelationships" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *XClusterReplicationFormData) contextValidateTargetAsyncReplicationRelationships(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "targetAsyncReplicationRelationships", "body", []*AsyncReplicationConfig(m.TargetAsyncReplicationRelationships)); err != nil {
		return err
	}

	for i := 0; i < len(m.TargetAsyncReplicationRelationships); i++ {

		if m.TargetAsyncReplicationRelationships[i] != nil {
			if err := m.TargetAsyncReplicationRelationships[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("targetAsyncReplicationRelationships" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *XClusterReplicationFormData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *XClusterReplicationFormData) UnmarshalBinary(b []byte) error {
	var res XClusterReplicationFormData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
