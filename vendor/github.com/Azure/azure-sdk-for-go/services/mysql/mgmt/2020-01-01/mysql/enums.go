package mysql

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// CreatedByType enumerates the values for created by type.
type CreatedByType string

const (
	// Application ...
	Application CreatedByType = "Application"
	// Key ...
	Key CreatedByType = "Key"
	// ManagedIdentity ...
	ManagedIdentity CreatedByType = "ManagedIdentity"
	// User ...
	User CreatedByType = "User"
)

// PossibleCreatedByTypeValues returns an array of possible values for the CreatedByType const type.
func PossibleCreatedByTypeValues() []CreatedByType {
	return []CreatedByType{Application, Key, ManagedIdentity, User}
}

// CreateMode enumerates the values for create mode.
type CreateMode string

const (
	// CreateModeDefault ...
	CreateModeDefault CreateMode = "Default"
	// CreateModeGeoRestore ...
	CreateModeGeoRestore CreateMode = "GeoRestore"
	// CreateModePointInTimeRestore ...
	CreateModePointInTimeRestore CreateMode = "PointInTimeRestore"
	// CreateModeReplica ...
	CreateModeReplica CreateMode = "Replica"
	// CreateModeServerPropertiesForCreate ...
	CreateModeServerPropertiesForCreate CreateMode = "ServerPropertiesForCreate"
)

// PossibleCreateModeValues returns an array of possible values for the CreateMode const type.
func PossibleCreateModeValues() []CreateMode {
	return []CreateMode{CreateModeDefault, CreateModeGeoRestore, CreateModePointInTimeRestore, CreateModeReplica, CreateModeServerPropertiesForCreate}
}

// GeoRedundantBackup enumerates the values for geo redundant backup.
type GeoRedundantBackup string

const (
	// Disabled ...
	Disabled GeoRedundantBackup = "Disabled"
	// Enabled ...
	Enabled GeoRedundantBackup = "Enabled"
)

// PossibleGeoRedundantBackupValues returns an array of possible values for the GeoRedundantBackup const type.
func PossibleGeoRedundantBackupValues() []GeoRedundantBackup {
	return []GeoRedundantBackup{Disabled, Enabled}
}

// IdentityType enumerates the values for identity type.
type IdentityType string

const (
	// SystemAssigned ...
	SystemAssigned IdentityType = "SystemAssigned"
)

// PossibleIdentityTypeValues returns an array of possible values for the IdentityType const type.
func PossibleIdentityTypeValues() []IdentityType {
	return []IdentityType{SystemAssigned}
}

// InfrastructureEncryption enumerates the values for infrastructure encryption.
type InfrastructureEncryption string

const (
	// InfrastructureEncryptionDisabled Additional (2nd) layer of encryption for data at rest
	InfrastructureEncryptionDisabled InfrastructureEncryption = "Disabled"
	// InfrastructureEncryptionEnabled Default value for single layer of encryption for data at rest.
	InfrastructureEncryptionEnabled InfrastructureEncryption = "Enabled"
)

// PossibleInfrastructureEncryptionValues returns an array of possible values for the InfrastructureEncryption const type.
func PossibleInfrastructureEncryptionValues() []InfrastructureEncryption {
	return []InfrastructureEncryption{InfrastructureEncryptionDisabled, InfrastructureEncryptionEnabled}
}

// MinimalTLSVersionEnum enumerates the values for minimal tls version enum.
type MinimalTLSVersionEnum string

const (
	// TLS10 ...
	TLS10 MinimalTLSVersionEnum = "TLS1_0"
	// TLS11 ...
	TLS11 MinimalTLSVersionEnum = "TLS1_1"
	// TLS12 ...
	TLS12 MinimalTLSVersionEnum = "TLS1_2"
	// TLSEnforcementDisabled ...
	TLSEnforcementDisabled MinimalTLSVersionEnum = "TLSEnforcementDisabled"
)

// PossibleMinimalTLSVersionEnumValues returns an array of possible values for the MinimalTLSVersionEnum const type.
func PossibleMinimalTLSVersionEnumValues() []MinimalTLSVersionEnum {
	return []MinimalTLSVersionEnum{TLS10, TLS11, TLS12, TLSEnforcementDisabled}
}

// OperationOrigin enumerates the values for operation origin.
type OperationOrigin string

const (
	// OperationOriginNotSpecified ...
	OperationOriginNotSpecified OperationOrigin = "NotSpecified"
	// OperationOriginSystem ...
	OperationOriginSystem OperationOrigin = "system"
	// OperationOriginUser ...
	OperationOriginUser OperationOrigin = "user"
)

// PossibleOperationOriginValues returns an array of possible values for the OperationOrigin const type.
func PossibleOperationOriginValues() []OperationOrigin {
	return []OperationOrigin{OperationOriginNotSpecified, OperationOriginSystem, OperationOriginUser}
}

// PrivateEndpointProvisioningState enumerates the values for private endpoint provisioning state.
type PrivateEndpointProvisioningState string

const (
	// Approving ...
	Approving PrivateEndpointProvisioningState = "Approving"
	// Dropping ...
	Dropping PrivateEndpointProvisioningState = "Dropping"
	// Failed ...
	Failed PrivateEndpointProvisioningState = "Failed"
	// Ready ...
	Ready PrivateEndpointProvisioningState = "Ready"
	// Rejecting ...
	Rejecting PrivateEndpointProvisioningState = "Rejecting"
)

// PossiblePrivateEndpointProvisioningStateValues returns an array of possible values for the PrivateEndpointProvisioningState const type.
func PossiblePrivateEndpointProvisioningStateValues() []PrivateEndpointProvisioningState {
	return []PrivateEndpointProvisioningState{Approving, Dropping, Failed, Ready, Rejecting}
}

// PrivateLinkServiceConnectionStateActionsRequire enumerates the values for private link service connection
// state actions require.
type PrivateLinkServiceConnectionStateActionsRequire string

const (
	// None ...
	None PrivateLinkServiceConnectionStateActionsRequire = "None"
)

// PossiblePrivateLinkServiceConnectionStateActionsRequireValues returns an array of possible values for the PrivateLinkServiceConnectionStateActionsRequire const type.
func PossiblePrivateLinkServiceConnectionStateActionsRequireValues() []PrivateLinkServiceConnectionStateActionsRequire {
	return []PrivateLinkServiceConnectionStateActionsRequire{None}
}

// PrivateLinkServiceConnectionStateStatus enumerates the values for private link service connection state
// status.
type PrivateLinkServiceConnectionStateStatus string

const (
	// Approved ...
	Approved PrivateLinkServiceConnectionStateStatus = "Approved"
	// Disconnected ...
	Disconnected PrivateLinkServiceConnectionStateStatus = "Disconnected"
	// Pending ...
	Pending PrivateLinkServiceConnectionStateStatus = "Pending"
	// Rejected ...
	Rejected PrivateLinkServiceConnectionStateStatus = "Rejected"
)

// PossiblePrivateLinkServiceConnectionStateStatusValues returns an array of possible values for the PrivateLinkServiceConnectionStateStatus const type.
func PossiblePrivateLinkServiceConnectionStateStatusValues() []PrivateLinkServiceConnectionStateStatus {
	return []PrivateLinkServiceConnectionStateStatus{Approved, Disconnected, Pending, Rejected}
}

// PublicNetworkAccessEnum enumerates the values for public network access enum.
type PublicNetworkAccessEnum string

const (
	// PublicNetworkAccessEnumDisabled ...
	PublicNetworkAccessEnumDisabled PublicNetworkAccessEnum = "Disabled"
	// PublicNetworkAccessEnumEnabled ...
	PublicNetworkAccessEnumEnabled PublicNetworkAccessEnum = "Enabled"
)

// PossiblePublicNetworkAccessEnumValues returns an array of possible values for the PublicNetworkAccessEnum const type.
func PossiblePublicNetworkAccessEnumValues() []PublicNetworkAccessEnum {
	return []PublicNetworkAccessEnum{PublicNetworkAccessEnumDisabled, PublicNetworkAccessEnumEnabled}
}

// ServerSecurityAlertPolicyState enumerates the values for server security alert policy state.
type ServerSecurityAlertPolicyState string

const (
	// ServerSecurityAlertPolicyStateDisabled ...
	ServerSecurityAlertPolicyStateDisabled ServerSecurityAlertPolicyState = "Disabled"
	// ServerSecurityAlertPolicyStateEnabled ...
	ServerSecurityAlertPolicyStateEnabled ServerSecurityAlertPolicyState = "Enabled"
)

// PossibleServerSecurityAlertPolicyStateValues returns an array of possible values for the ServerSecurityAlertPolicyState const type.
func PossibleServerSecurityAlertPolicyStateValues() []ServerSecurityAlertPolicyState {
	return []ServerSecurityAlertPolicyState{ServerSecurityAlertPolicyStateDisabled, ServerSecurityAlertPolicyStateEnabled}
}

// ServerState enumerates the values for server state.
type ServerState string

const (
	// ServerStateDisabled ...
	ServerStateDisabled ServerState = "Disabled"
	// ServerStateDropping ...
	ServerStateDropping ServerState = "Dropping"
	// ServerStateInaccessible ...
	ServerStateInaccessible ServerState = "Inaccessible"
	// ServerStateReady ...
	ServerStateReady ServerState = "Ready"
)

// PossibleServerStateValues returns an array of possible values for the ServerState const type.
func PossibleServerStateValues() []ServerState {
	return []ServerState{ServerStateDisabled, ServerStateDropping, ServerStateInaccessible, ServerStateReady}
}

// ServerVersion enumerates the values for server version.
type ServerVersion string

const (
	// EightFullStopZero ...
	EightFullStopZero ServerVersion = "8.0"
	// FiveFullStopSeven ...
	FiveFullStopSeven ServerVersion = "5.7"
	// FiveFullStopSix ...
	FiveFullStopSix ServerVersion = "5.6"
)

// PossibleServerVersionValues returns an array of possible values for the ServerVersion const type.
func PossibleServerVersionValues() []ServerVersion {
	return []ServerVersion{EightFullStopZero, FiveFullStopSeven, FiveFullStopSix}
}

// SkuTier enumerates the values for sku tier.
type SkuTier string

const (
	// Basic ...
	Basic SkuTier = "Basic"
	// GeneralPurpose ...
	GeneralPurpose SkuTier = "GeneralPurpose"
	// MemoryOptimized ...
	MemoryOptimized SkuTier = "MemoryOptimized"
)

// PossibleSkuTierValues returns an array of possible values for the SkuTier const type.
func PossibleSkuTierValues() []SkuTier {
	return []SkuTier{Basic, GeneralPurpose, MemoryOptimized}
}

// SslEnforcementEnum enumerates the values for ssl enforcement enum.
type SslEnforcementEnum string

const (
	// SslEnforcementEnumDisabled ...
	SslEnforcementEnumDisabled SslEnforcementEnum = "Disabled"
	// SslEnforcementEnumEnabled ...
	SslEnforcementEnumEnabled SslEnforcementEnum = "Enabled"
)

// PossibleSslEnforcementEnumValues returns an array of possible values for the SslEnforcementEnum const type.
func PossibleSslEnforcementEnumValues() []SslEnforcementEnum {
	return []SslEnforcementEnum{SslEnforcementEnumDisabled, SslEnforcementEnumEnabled}
}

// StorageAutogrow enumerates the values for storage autogrow.
type StorageAutogrow string

const (
	// StorageAutogrowDisabled ...
	StorageAutogrowDisabled StorageAutogrow = "Disabled"
	// StorageAutogrowEnabled ...
	StorageAutogrowEnabled StorageAutogrow = "Enabled"
)

// PossibleStorageAutogrowValues returns an array of possible values for the StorageAutogrow const type.
func PossibleStorageAutogrowValues() []StorageAutogrow {
	return []StorageAutogrow{StorageAutogrowDisabled, StorageAutogrowEnabled}
}

// VirtualNetworkRuleState enumerates the values for virtual network rule state.
type VirtualNetworkRuleState string

const (
	// VirtualNetworkRuleStateDeleting ...
	VirtualNetworkRuleStateDeleting VirtualNetworkRuleState = "Deleting"
	// VirtualNetworkRuleStateInitializing ...
	VirtualNetworkRuleStateInitializing VirtualNetworkRuleState = "Initializing"
	// VirtualNetworkRuleStateInProgress ...
	VirtualNetworkRuleStateInProgress VirtualNetworkRuleState = "InProgress"
	// VirtualNetworkRuleStateReady ...
	VirtualNetworkRuleStateReady VirtualNetworkRuleState = "Ready"
	// VirtualNetworkRuleStateUnknown ...
	VirtualNetworkRuleStateUnknown VirtualNetworkRuleState = "Unknown"
)

// PossibleVirtualNetworkRuleStateValues returns an array of possible values for the VirtualNetworkRuleState const type.
func PossibleVirtualNetworkRuleStateValues() []VirtualNetworkRuleState {
	return []VirtualNetworkRuleState{VirtualNetworkRuleStateDeleting, VirtualNetworkRuleStateInitializing, VirtualNetworkRuleStateInProgress, VirtualNetworkRuleStateReady, VirtualNetworkRuleStateUnknown}
}
