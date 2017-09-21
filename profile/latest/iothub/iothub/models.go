// +build go1.9

// Copyright 2017 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder
// commit ID: b19c730e3a5c747d9055c95884b5c0310f7f2f16

package iothub

import original "github.com/Azure/azure-sdk-for-go/service/iothub/management/2016-02-03/iothub"

type ResourceClient = original.ResourceClient

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ManagementClient = original.ManagementClient
type AccessRights = original.AccessRights

const (
	DeviceConnect						AccessRights	= original.DeviceConnect
	RegistryRead						AccessRights	= original.RegistryRead
	RegistryReadDeviceConnect				AccessRights	= original.RegistryReadDeviceConnect
	RegistryReadRegistryWrite				AccessRights	= original.RegistryReadRegistryWrite
	RegistryReadRegistryWriteDeviceConnect			AccessRights	= original.RegistryReadRegistryWriteDeviceConnect
	RegistryReadRegistryWriteServiceConnect			AccessRights	= original.RegistryReadRegistryWriteServiceConnect
	RegistryReadRegistryWriteServiceConnectDeviceConnect	AccessRights	= original.RegistryReadRegistryWriteServiceConnectDeviceConnect
	RegistryReadServiceConnect				AccessRights	= original.RegistryReadServiceConnect
	RegistryReadServiceConnectDeviceConnect			AccessRights	= original.RegistryReadServiceConnectDeviceConnect
	RegistryWrite						AccessRights	= original.RegistryWrite
	RegistryWriteDeviceConnect				AccessRights	= original.RegistryWriteDeviceConnect
	RegistryWriteServiceConnect				AccessRights	= original.RegistryWriteServiceConnect
	RegistryWriteServiceConnectDeviceConnect		AccessRights	= original.RegistryWriteServiceConnectDeviceConnect
	ServiceConnect						AccessRights	= original.ServiceConnect
	ServiceConnectDeviceConnect				AccessRights	= original.ServiceConnectDeviceConnect
)

type Capabilities = original.Capabilities

const (
	DeviceManagement	Capabilities	= original.DeviceManagement
	None			Capabilities	= original.None
)

type IPFilterActionType = original.IPFilterActionType

const (
	Accept	IPFilterActionType	= original.Accept
	Reject	IPFilterActionType	= original.Reject
)

type JobStatus = original.JobStatus

const (
	Cancelled	JobStatus	= original.Cancelled
	Completed	JobStatus	= original.Completed
	Enqueued	JobStatus	= original.Enqueued
	Failed		JobStatus	= original.Failed
	Running		JobStatus	= original.Running
	Unknown		JobStatus	= original.Unknown
)

type JobType = original.JobType

const (
	JobTypeBackup				JobType	= original.JobTypeBackup
	JobTypeExport				JobType	= original.JobTypeExport
	JobTypeFactoryResetDevice		JobType	= original.JobTypeFactoryResetDevice
	JobTypeFirmwareUpdate			JobType	= original.JobTypeFirmwareUpdate
	JobTypeImport				JobType	= original.JobTypeImport
	JobTypeReadDeviceProperties		JobType	= original.JobTypeReadDeviceProperties
	JobTypeRebootDevice			JobType	= original.JobTypeRebootDevice
	JobTypeUnknown				JobType	= original.JobTypeUnknown
	JobTypeUpdateDeviceConfiguration	JobType	= original.JobTypeUpdateDeviceConfiguration
	JobTypeWriteDeviceProperties		JobType	= original.JobTypeWriteDeviceProperties
)

type NameUnavailabilityReason = original.NameUnavailabilityReason

const (
	AlreadyExists	NameUnavailabilityReason	= original.AlreadyExists
	Invalid		NameUnavailabilityReason	= original.Invalid
)

type OperationMonitoringLevel = original.OperationMonitoringLevel

const (
	OperationMonitoringLevelError			OperationMonitoringLevel	= original.OperationMonitoringLevelError
	OperationMonitoringLevelErrorInformation	OperationMonitoringLevel	= original.OperationMonitoringLevelErrorInformation
	OperationMonitoringLevelInformation		OperationMonitoringLevel	= original.OperationMonitoringLevelInformation
	OperationMonitoringLevelNone			OperationMonitoringLevel	= original.OperationMonitoringLevelNone
)

type ScaleType = original.ScaleType

const (
	ScaleTypeAutomatic	ScaleType	= original.ScaleTypeAutomatic
	ScaleTypeManual		ScaleType	= original.ScaleTypeManual
	ScaleTypeNone		ScaleType	= original.ScaleTypeNone
)

type Sku = original.Sku

const (
	F1	Sku	= original.F1
	S1	Sku	= original.S1
	S2	Sku	= original.S2
	S3	Sku	= original.S3
)

type SkuTier = original.SkuTier

const (
	Free		SkuTier	= original.Free
	Standard	SkuTier	= original.Standard
)

type Capacity = original.Capacity
type CloudToDeviceProperties = original.CloudToDeviceProperties
type Description = original.Description
type DescriptionListResult = original.DescriptionListResult
type ErrorDetails = original.ErrorDetails
type EventHubConsumerGroupInfo = original.EventHubConsumerGroupInfo
type EventHubConsumerGroupsListResult = original.EventHubConsumerGroupsListResult
type EventHubProperties = original.EventHubProperties
type ExportDevicesRequest = original.ExportDevicesRequest
type FeedbackProperties = original.FeedbackProperties
type ImportDevicesRequest = original.ImportDevicesRequest
type IPFilterRule = original.IPFilterRule
type JobResponse = original.JobResponse
type JobResponseListResult = original.JobResponseListResult
type MessagingEndpointProperties = original.MessagingEndpointProperties
type NameAvailabilityInfo = original.NameAvailabilityInfo
type OperationInputs = original.OperationInputs
type OperationsMonitoringProperties = original.OperationsMonitoringProperties
type Properties = original.Properties
type QuotaMetricInfo = original.QuotaMetricInfo
type QuotaMetricInfoListResult = original.QuotaMetricInfoListResult
type RegistryStatistics = original.RegistryStatistics
type Resource = original.Resource
type SetObject = original.SetObject
type SharedAccessSignatureAuthorizationRule = original.SharedAccessSignatureAuthorizationRule
type SharedAccessSignatureAuthorizationRuleListResult = original.SharedAccessSignatureAuthorizationRuleListResult
type SkuDescription = original.SkuDescription
type SkuDescriptionListResult = original.SkuDescriptionListResult
type SkuInfo = original.SkuInfo
type StorageEndpointProperties = original.StorageEndpointProperties

func New(subscriptionID string) ManagementClient {
	return original.New(subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) ManagementClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func NewResourceClient(subscriptionID string) ResourceClient {
	return original.NewResourceClient(subscriptionID)
}
func NewResourceClientWithBaseURI(baseURI string, subscriptionID string) ResourceClient {
	return original.NewResourceClientWithBaseURI(baseURI, subscriptionID)
}
func UserAgent() string {
	return original.UserAgent()
}
func Version() string {
	return original.Version()
}