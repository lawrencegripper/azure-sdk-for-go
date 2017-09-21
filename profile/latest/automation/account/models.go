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

package account

import original "github.com/Azure/azure-sdk-for-go/service/automation/management/2015-10-31/account"

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ManagementClient = original.ManagementClient
type AutomationAccountState = original.AutomationAccountState

const (
	Ok		AutomationAccountState	= original.Ok
	Suspended	AutomationAccountState	= original.Suspended
	Unavailable	AutomationAccountState	= original.Unavailable
)

type ContentSourceType = original.ContentSourceType

const (
	EmbeddedContent	ContentSourceType	= original.EmbeddedContent
	URI		ContentSourceType	= original.URI
)

type DscConfigurationProvisioningState = original.DscConfigurationProvisioningState

const (
	Succeeded DscConfigurationProvisioningState = original.Succeeded
)

type DscConfigurationState = original.DscConfigurationState

const (
	DscConfigurationStateEdit	DscConfigurationState	= original.DscConfigurationStateEdit
	DscConfigurationStateNew	DscConfigurationState	= original.DscConfigurationStateNew
	DscConfigurationStatePublished	DscConfigurationState	= original.DscConfigurationStatePublished
)

type ModuleProvisioningState = original.ModuleProvisioningState

const (
	ModuleProvisioningStateActivitiesStored			ModuleProvisioningState	= original.ModuleProvisioningStateActivitiesStored
	ModuleProvisioningStateCancelled			ModuleProvisioningState	= original.ModuleProvisioningStateCancelled
	ModuleProvisioningStateConnectionTypeImported		ModuleProvisioningState	= original.ModuleProvisioningStateConnectionTypeImported
	ModuleProvisioningStateContentDownloaded		ModuleProvisioningState	= original.ModuleProvisioningStateContentDownloaded
	ModuleProvisioningStateContentRetrieved			ModuleProvisioningState	= original.ModuleProvisioningStateContentRetrieved
	ModuleProvisioningStateContentStored			ModuleProvisioningState	= original.ModuleProvisioningStateContentStored
	ModuleProvisioningStateContentValidated			ModuleProvisioningState	= original.ModuleProvisioningStateContentValidated
	ModuleProvisioningStateCreated				ModuleProvisioningState	= original.ModuleProvisioningStateCreated
	ModuleProvisioningStateCreating				ModuleProvisioningState	= original.ModuleProvisioningStateCreating
	ModuleProvisioningStateFailed				ModuleProvisioningState	= original.ModuleProvisioningStateFailed
	ModuleProvisioningStateModuleDataStored			ModuleProvisioningState	= original.ModuleProvisioningStateModuleDataStored
	ModuleProvisioningStateModuleImportRunbookComplete	ModuleProvisioningState	= original.ModuleProvisioningStateModuleImportRunbookComplete
	ModuleProvisioningStateRunningImportModuleRunbook	ModuleProvisioningState	= original.ModuleProvisioningStateRunningImportModuleRunbook
	ModuleProvisioningStateStartingImportModuleRunbook	ModuleProvisioningState	= original.ModuleProvisioningStateStartingImportModuleRunbook
	ModuleProvisioningStateSucceeded			ModuleProvisioningState	= original.ModuleProvisioningStateSucceeded
	ModuleProvisioningStateUpdating				ModuleProvisioningState	= original.ModuleProvisioningStateUpdating
)

type RunbookProvisioningState = original.RunbookProvisioningState

const (
	RunbookProvisioningStateSucceeded RunbookProvisioningState = original.RunbookProvisioningStateSucceeded
)

type RunbookState = original.RunbookState

const (
	RunbookStateEdit	RunbookState	= original.RunbookStateEdit
	RunbookStateNew		RunbookState	= original.RunbookStateNew
	RunbookStatePublished	RunbookState	= original.RunbookStatePublished
)

type RunbookTypeEnum = original.RunbookTypeEnum

const (
	Graph			RunbookTypeEnum	= original.Graph
	GraphPowerShell		RunbookTypeEnum	= original.GraphPowerShell
	GraphPowerShellWorkflow	RunbookTypeEnum	= original.GraphPowerShellWorkflow
	PowerShell		RunbookTypeEnum	= original.PowerShell
	PowerShellWorkflow	RunbookTypeEnum	= original.PowerShellWorkflow
	Script			RunbookTypeEnum	= original.Script
)

type SkuNameEnum = original.SkuNameEnum

const (
	Basic	SkuNameEnum	= original.Basic
	Free	SkuNameEnum	= original.Free
)

type AutomationAccount = original.AutomationAccount
type AutomationAccountCreateOrUpdateParameters = original.AutomationAccountCreateOrUpdateParameters
type AutomationAccountCreateOrUpdateProperties = original.AutomationAccountCreateOrUpdateProperties
type AutomationAccountListResult = original.AutomationAccountListResult
type AutomationAccountProperties = original.AutomationAccountProperties
type AutomationAccountUpdateParameters = original.AutomationAccountUpdateParameters
type AutomationAccountUpdateProperties = original.AutomationAccountUpdateProperties
type ContentHash = original.ContentHash
type ContentLink = original.ContentLink
type ContentSource = original.ContentSource
type DscConfiguration = original.DscConfiguration
type DscConfigurationParameter = original.DscConfigurationParameter
type DscConfigurationProperties = original.DscConfigurationProperties
type DscNode = original.DscNode
type DscNodeConfigurationAssociationProperty = original.DscNodeConfigurationAssociationProperty
type ErrorResponse = original.ErrorResponse
type Module = original.Module
type ModuleErrorInfo = original.ModuleErrorInfo
type ModuleProperties = original.ModuleProperties
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type Resource = original.Resource
type Runbook = original.Runbook
type RunbookDraft = original.RunbookDraft
type RunbookParameter = original.RunbookParameter
type RunbookProperties = original.RunbookProperties
type Sku = original.Sku
type Statistics = original.Statistics
type StatisticsListResult = original.StatisticsListResult
type Usage = original.Usage
type UsageCounterName = original.UsageCounterName
type UsageListResult = original.UsageListResult
type OperationsClient = original.OperationsClient
type StatisticsClient = original.StatisticsClient
type UsagesClient = original.UsagesClient
type AutomationAccountClient = original.AutomationAccountClient

func New(subscriptionID string) ManagementClient {
	return original.New(subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) ManagementClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewStatisticsClient(subscriptionID string) StatisticsClient {
	return original.NewStatisticsClient(subscriptionID)
}
func NewStatisticsClientWithBaseURI(baseURI string, subscriptionID string) StatisticsClient {
	return original.NewStatisticsClientWithBaseURI(baseURI, subscriptionID)
}
func NewUsagesClient(subscriptionID string) UsagesClient {
	return original.NewUsagesClient(subscriptionID)
}
func NewUsagesClientWithBaseURI(baseURI string, subscriptionID string) UsagesClient {
	return original.NewUsagesClientWithBaseURI(baseURI, subscriptionID)
}
func UserAgent() string {
	return original.UserAgent()
}
func Version() string {
	return original.Version()
}
func NewAutomationAccountClient(subscriptionID string) AutomationAccountClient {
	return original.NewAutomationAccountClient(subscriptionID)
}
func NewAutomationAccountClientWithBaseURI(baseURI string, subscriptionID string) AutomationAccountClient {
	return original.NewAutomationAccountClientWithBaseURI(baseURI, subscriptionID)
}
