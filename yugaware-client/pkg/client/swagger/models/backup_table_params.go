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

// BackupTableParams Backup table parameters
//
// swagger:model BackupTableParams
type BackupTableParams struct {

	// Action type
	// Required: true
	// Enum: [CREATE RESTORE RESTORE_KEYS DELETE]
	ActionType *string `json:"actionType"`

	// Backups
	BackupList []*BackupTableParams `json:"backupList"`

	// Backup type
	// Enum: [YQL_TABLE_TYPE REDIS_TABLE_TYPE PGSQL_TABLE_TYPE TRANSACTION_STATUS_TABLE_TYPE]
	BackupType string `json:"backupType,omitempty"`

	// Backup UUID
	// Format: uuid
	BackupUUID strfmt.UUID `json:"backupUuid,omitempty"`

	// Amazon Resource Name (ARN) of the CMK
	CmkArn string `json:"cmkArn,omitempty"`

	// Communication ports
	CommunicationPorts *CommunicationPorts `json:"communicationPorts,omitempty"`

	// Cron expression for a recurring backup
	CronExpression string `json:"cronExpression,omitempty"`

	// Customer UUID
	// Format: uuid
	CustomerUUID strfmt.UUID `json:"customerUUID,omitempty"`

	// Device information
	DeviceInfo *DeviceInfo `json:"deviceInfo,omitempty"`

	// Is verbose logging enabled
	EnableVerboseLogs bool `json:"enableVerboseLogs,omitempty"`

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

	// Should table backup errors be ignored
	IgnoreErrors bool `json:"ignoreErrors,omitempty"`

	// Key space
	Keyspace string `json:"keyspace,omitempty"`

	// KMS configuration UUID
	// Format: uuid
	KmsConfigUUID strfmt.UUID `json:"kmsConfigUUID,omitempty"`

	// Node details
	// Unique: true
	NodeDetailsSet []*NodeDetails `json:"nodeDetailsSet"`

	// Node exporter user
	NodeExporterUser string `json:"nodeExporterUser,omitempty"`

	// Number of concurrent commands to run on nodes over SSH
	Parallelism int32 `json:"parallelism,omitempty"`

	// Restore TimeStamp
	RestoreTimeStamp string `json:"restoreTimeStamp,omitempty"`

	// Schedule UUID
	// Format: uuid
	ScheduleUUID strfmt.UUID `json:"scheduleUUID,omitempty"`

	// Frequency to run the backup, in milliseconds
	SchedulingFrequency int64 `json:"schedulingFrequency,omitempty"`

	// The source universe's sync replication relationships
	// Read Only: true
	SourceAsyncReplicationRelationships []*AsyncReplicationConfig `json:"sourceAsyncReplicationRelationships"`

	// Is SSE
	Sse bool `json:"sse,omitempty"`

	// Storage configuration UUID
	// Required: true
	// Format: uuid
	StorageConfigUUID *strfmt.UUID `json:"storageConfigUUID"`

	// Storage location
	StorageLocation string `json:"storageLocation,omitempty"`

	// Table name
	TableName string `json:"tableName,omitempty"`

	// Tables
	TableNameList []string `json:"tableNameList"`

	// Table UUID
	// Format: uuid
	TableUUID strfmt.UUID `json:"tableUUID,omitempty"`

	// Table UUIDs
	TableUUIDList []strfmt.UUID `json:"tableUUIDList"`

	// The target universe's async replication relationships
	// Read Only: true
	TargetAsyncReplicationRelationships []*AsyncReplicationConfig `json:"targetAsyncReplicationRelationships"`

	// Time before deleting the backup from storage, in milliseconds
	TimeBeforeDelete int64 `json:"timeBeforeDelete,omitempty"`

	// Is backup transactional across tables
	TransactionalBackup bool `json:"transactionalBackup,omitempty"`

	// Associated universe UUID
	// Format: uuid
	UniverseUUID strfmt.UUID `json:"universeUUID,omitempty"`

	// Previous software version
	YbPrevSoftwareVersion string `json:"ybPrevSoftwareVersion,omitempty"`
}

// Validate validates this backup table params
func (m *BackupTableParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActionType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBackupList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBackupType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBackupUUID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCommunicationPorts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCustomerUUID(formats); err != nil {
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

	if err := m.validateKmsConfigUUID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNodeDetailsSet(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScheduleUUID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceAsyncReplicationRelationships(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStorageConfigUUID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTableUUID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTableUUIDList(formats); err != nil {
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

var backupTableParamsTypeActionTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CREATE","RESTORE","RESTORE_KEYS","DELETE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		backupTableParamsTypeActionTypePropEnum = append(backupTableParamsTypeActionTypePropEnum, v)
	}
}

const (

	// BackupTableParamsActionTypeCREATE captures enum value "CREATE"
	BackupTableParamsActionTypeCREATE string = "CREATE"

	// BackupTableParamsActionTypeRESTORE captures enum value "RESTORE"
	BackupTableParamsActionTypeRESTORE string = "RESTORE"

	// BackupTableParamsActionTypeRESTOREKEYS captures enum value "RESTORE_KEYS"
	BackupTableParamsActionTypeRESTOREKEYS string = "RESTORE_KEYS"

	// BackupTableParamsActionTypeDELETE captures enum value "DELETE"
	BackupTableParamsActionTypeDELETE string = "DELETE"
)

// prop value enum
func (m *BackupTableParams) validateActionTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, backupTableParamsTypeActionTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *BackupTableParams) validateActionType(formats strfmt.Registry) error {

	if err := validate.Required("actionType", "body", m.ActionType); err != nil {
		return err
	}

	// value enum
	if err := m.validateActionTypeEnum("actionType", "body", *m.ActionType); err != nil {
		return err
	}

	return nil
}

func (m *BackupTableParams) validateBackupList(formats strfmt.Registry) error {
	if swag.IsZero(m.BackupList) { // not required
		return nil
	}

	for i := 0; i < len(m.BackupList); i++ {
		if swag.IsZero(m.BackupList[i]) { // not required
			continue
		}

		if m.BackupList[i] != nil {
			if err := m.BackupList[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("backupList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var backupTableParamsTypeBackupTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["YQL_TABLE_TYPE","REDIS_TABLE_TYPE","PGSQL_TABLE_TYPE","TRANSACTION_STATUS_TABLE_TYPE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		backupTableParamsTypeBackupTypePropEnum = append(backupTableParamsTypeBackupTypePropEnum, v)
	}
}

const (

	// BackupTableParamsBackupTypeYQLTABLETYPE captures enum value "YQL_TABLE_TYPE"
	BackupTableParamsBackupTypeYQLTABLETYPE string = "YQL_TABLE_TYPE"

	// BackupTableParamsBackupTypeREDISTABLETYPE captures enum value "REDIS_TABLE_TYPE"
	BackupTableParamsBackupTypeREDISTABLETYPE string = "REDIS_TABLE_TYPE"

	// BackupTableParamsBackupTypePGSQLTABLETYPE captures enum value "PGSQL_TABLE_TYPE"
	BackupTableParamsBackupTypePGSQLTABLETYPE string = "PGSQL_TABLE_TYPE"

	// BackupTableParamsBackupTypeTRANSACTIONSTATUSTABLETYPE captures enum value "TRANSACTION_STATUS_TABLE_TYPE"
	BackupTableParamsBackupTypeTRANSACTIONSTATUSTABLETYPE string = "TRANSACTION_STATUS_TABLE_TYPE"
)

// prop value enum
func (m *BackupTableParams) validateBackupTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, backupTableParamsTypeBackupTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *BackupTableParams) validateBackupType(formats strfmt.Registry) error {
	if swag.IsZero(m.BackupType) { // not required
		return nil
	}

	// value enum
	if err := m.validateBackupTypeEnum("backupType", "body", m.BackupType); err != nil {
		return err
	}

	return nil
}

func (m *BackupTableParams) validateBackupUUID(formats strfmt.Registry) error {
	if swag.IsZero(m.BackupUUID) { // not required
		return nil
	}

	if err := validate.FormatOf("backupUuid", "body", "uuid", m.BackupUUID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *BackupTableParams) validateCommunicationPorts(formats strfmt.Registry) error {
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

func (m *BackupTableParams) validateCustomerUUID(formats strfmt.Registry) error {
	if swag.IsZero(m.CustomerUUID) { // not required
		return nil
	}

	if err := validate.FormatOf("customerUUID", "body", "uuid", m.CustomerUUID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *BackupTableParams) validateDeviceInfo(formats strfmt.Registry) error {
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

func (m *BackupTableParams) validateEncryptionAtRestConfig(formats strfmt.Registry) error {
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

func (m *BackupTableParams) validateExtraDependencies(formats strfmt.Registry) error {
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

func (m *BackupTableParams) validateKmsConfigUUID(formats strfmt.Registry) error {
	if swag.IsZero(m.KmsConfigUUID) { // not required
		return nil
	}

	if err := validate.FormatOf("kmsConfigUUID", "body", "uuid", m.KmsConfigUUID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *BackupTableParams) validateNodeDetailsSet(formats strfmt.Registry) error {
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

func (m *BackupTableParams) validateScheduleUUID(formats strfmt.Registry) error {
	if swag.IsZero(m.ScheduleUUID) { // not required
		return nil
	}

	if err := validate.FormatOf("scheduleUUID", "body", "uuid", m.ScheduleUUID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *BackupTableParams) validateSourceAsyncReplicationRelationships(formats strfmt.Registry) error {
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

func (m *BackupTableParams) validateStorageConfigUUID(formats strfmt.Registry) error {

	if err := validate.Required("storageConfigUUID", "body", m.StorageConfigUUID); err != nil {
		return err
	}

	if err := validate.FormatOf("storageConfigUUID", "body", "uuid", m.StorageConfigUUID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *BackupTableParams) validateTableUUID(formats strfmt.Registry) error {
	if swag.IsZero(m.TableUUID) { // not required
		return nil
	}

	if err := validate.FormatOf("tableUUID", "body", "uuid", m.TableUUID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *BackupTableParams) validateTableUUIDList(formats strfmt.Registry) error {
	if swag.IsZero(m.TableUUIDList) { // not required
		return nil
	}

	for i := 0; i < len(m.TableUUIDList); i++ {

		if err := validate.FormatOf("tableUUIDList"+"."+strconv.Itoa(i), "body", "uuid", m.TableUUIDList[i].String(), formats); err != nil {
			return err
		}

	}

	return nil
}

func (m *BackupTableParams) validateTargetAsyncReplicationRelationships(formats strfmt.Registry) error {
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

func (m *BackupTableParams) validateUniverseUUID(formats strfmt.Registry) error {
	if swag.IsZero(m.UniverseUUID) { // not required
		return nil
	}

	if err := validate.FormatOf("universeUUID", "body", "uuid", m.UniverseUUID.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this backup table params based on the context it is used
func (m *BackupTableParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBackupList(ctx, formats); err != nil {
		res = append(res, err)
	}

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

func (m *BackupTableParams) contextValidateBackupList(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.BackupList); i++ {

		if m.BackupList[i] != nil {
			if err := m.BackupList[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("backupList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BackupTableParams) contextValidateCommunicationPorts(ctx context.Context, formats strfmt.Registry) error {

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

func (m *BackupTableParams) contextValidateDeviceInfo(ctx context.Context, formats strfmt.Registry) error {

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

func (m *BackupTableParams) contextValidateEncryptionAtRestConfig(ctx context.Context, formats strfmt.Registry) error {

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

func (m *BackupTableParams) contextValidateExtraDependencies(ctx context.Context, formats strfmt.Registry) error {

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

func (m *BackupTableParams) contextValidateNodeDetailsSet(ctx context.Context, formats strfmt.Registry) error {

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

func (m *BackupTableParams) contextValidateSourceAsyncReplicationRelationships(ctx context.Context, formats strfmt.Registry) error {

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

func (m *BackupTableParams) contextValidateTargetAsyncReplicationRelationships(ctx context.Context, formats strfmt.Registry) error {

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
func (m *BackupTableParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BackupTableParams) UnmarshalBinary(b []byte) error {
	var res BackupTableParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
