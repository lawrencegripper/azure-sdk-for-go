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

import original "github.com/Azure/azure-sdk-for-go/service/datalake-analytics/management/2016-11-01/account"

type GroupClient = original.GroupClient

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ManagementClient = original.ManagementClient
type ComputePoliciesClient = original.ComputePoliciesClient
type DataLakeStoreAccountsClient = original.DataLakeStoreAccountsClient
type FirewallRulesClient = original.FirewallRulesClient
type AADObjectType = original.AADObjectType

const (
	Group			AADObjectType	= original.Group
	ServicePrincipal	AADObjectType	= original.ServicePrincipal
	User			AADObjectType	= original.User
)

type DataLakeAnalyticsAccountState = original.DataLakeAnalyticsAccountState

const (
	Active		DataLakeAnalyticsAccountState	= original.Active
	Suspended	DataLakeAnalyticsAccountState	= original.Suspended
)

type DataLakeAnalyticsAccountStatus = original.DataLakeAnalyticsAccountStatus

const (
	Creating	DataLakeAnalyticsAccountStatus	= original.Creating
	Deleted		DataLakeAnalyticsAccountStatus	= original.Deleted
	Deleting	DataLakeAnalyticsAccountStatus	= original.Deleting
	Failed		DataLakeAnalyticsAccountStatus	= original.Failed
	Patching	DataLakeAnalyticsAccountStatus	= original.Patching
	Resuming	DataLakeAnalyticsAccountStatus	= original.Resuming
	Running		DataLakeAnalyticsAccountStatus	= original.Running
	Succeeded	DataLakeAnalyticsAccountStatus	= original.Succeeded
	Suspending	DataLakeAnalyticsAccountStatus	= original.Suspending
)

type FirewallAllowAzureIpsState = original.FirewallAllowAzureIpsState

const (
	Disabled	FirewallAllowAzureIpsState	= original.Disabled
	Enabled		FirewallAllowAzureIpsState	= original.Enabled
)

type FirewallState = original.FirewallState

const (
	FirewallStateDisabled	FirewallState	= original.FirewallStateDisabled
	FirewallStateEnabled	FirewallState	= original.FirewallStateEnabled
)

type TierType = original.TierType

const (
	Commitment100000AUHours	TierType	= original.Commitment100000AUHours
	Commitment10000AUHours	TierType	= original.Commitment10000AUHours
	Commitment1000AUHours	TierType	= original.Commitment1000AUHours
	Commitment100AUHours	TierType	= original.Commitment100AUHours
	Commitment500000AUHours	TierType	= original.Commitment500000AUHours
	Commitment50000AUHours	TierType	= original.Commitment50000AUHours
	Commitment5000AUHours	TierType	= original.Commitment5000AUHours
	Commitment500AUHours	TierType	= original.Commitment500AUHours
	Consumption		TierType	= original.Consumption
)

type AddDataLakeStoreParameters = original.AddDataLakeStoreParameters
type AddStorageAccountParameters = original.AddStorageAccountParameters
type ComputePolicy = original.ComputePolicy
type ComputePolicyAccountCreateParameters = original.ComputePolicyAccountCreateParameters
type ComputePolicyCreateOrUpdateParameters = original.ComputePolicyCreateOrUpdateParameters
type ComputePolicyListResult = original.ComputePolicyListResult
type ComputePolicyProperties = original.ComputePolicyProperties
type ComputePolicyPropertiesCreateParameters = original.ComputePolicyPropertiesCreateParameters
type DataLakeAnalyticsAccount = original.DataLakeAnalyticsAccount
type DataLakeAnalyticsAccountBasic = original.DataLakeAnalyticsAccountBasic
type DataLakeAnalyticsAccountListDataLakeStoreResult = original.DataLakeAnalyticsAccountListDataLakeStoreResult
type DataLakeAnalyticsAccountListResult = original.DataLakeAnalyticsAccountListResult
type DataLakeAnalyticsAccountListStorageAccountsResult = original.DataLakeAnalyticsAccountListStorageAccountsResult
type DataLakeAnalyticsAccountProperties = original.DataLakeAnalyticsAccountProperties
type DataLakeAnalyticsAccountPropertiesBasic = original.DataLakeAnalyticsAccountPropertiesBasic
type DataLakeAnalyticsAccountUpdateParameters = original.DataLakeAnalyticsAccountUpdateParameters
type DataLakeAnalyticsFirewallRuleListResult = original.DataLakeAnalyticsFirewallRuleListResult
type DataLakeStoreAccountInfo = original.DataLakeStoreAccountInfo
type DataLakeStoreAccountInfoProperties = original.DataLakeStoreAccountInfoProperties
type FirewallRule = original.FirewallRule
type FirewallRuleProperties = original.FirewallRuleProperties
type ListSasTokensResult = original.ListSasTokensResult
type ListStorageContainersResult = original.ListStorageContainersResult
type OptionalSubResource = original.OptionalSubResource
type Resource = original.Resource
type SasTokenInfo = original.SasTokenInfo
type StorageAccountInfo = original.StorageAccountInfo
type StorageAccountProperties = original.StorageAccountProperties
type StorageContainer = original.StorageContainer
type StorageContainerProperties = original.StorageContainerProperties
type SubResource = original.SubResource
type UpdateDataLakeAnalyticsAccountProperties = original.UpdateDataLakeAnalyticsAccountProperties
type UpdateFirewallRuleParameters = original.UpdateFirewallRuleParameters
type UpdateFirewallRuleProperties = original.UpdateFirewallRuleProperties
type UpdateStorageAccountParameters = original.UpdateStorageAccountParameters
type UpdateStorageAccountProperties = original.UpdateStorageAccountProperties
type StorageAccountsClient = original.StorageAccountsClient

func NewFirewallRulesClient(subscriptionID string) FirewallRulesClient {
	return original.NewFirewallRulesClient(subscriptionID)
}
func NewFirewallRulesClientWithBaseURI(baseURI string, subscriptionID string) FirewallRulesClient {
	return original.NewFirewallRulesClientWithBaseURI(baseURI, subscriptionID)
}
func NewStorageAccountsClient(subscriptionID string) StorageAccountsClient {
	return original.NewStorageAccountsClient(subscriptionID)
}
func NewStorageAccountsClientWithBaseURI(baseURI string, subscriptionID string) StorageAccountsClient {
	return original.NewStorageAccountsClientWithBaseURI(baseURI, subscriptionID)
}
func UserAgent() string {
	return original.UserAgent()
}
func Version() string {
	return original.Version()
}
func NewGroupClient(subscriptionID string) GroupClient {
	return original.NewGroupClient(subscriptionID)
}
func NewGroupClientWithBaseURI(baseURI string, subscriptionID string) GroupClient {
	return original.NewGroupClientWithBaseURI(baseURI, subscriptionID)
}
func New(subscriptionID string) ManagementClient {
	return original.New(subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) ManagementClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func NewComputePoliciesClient(subscriptionID string) ComputePoliciesClient {
	return original.NewComputePoliciesClient(subscriptionID)
}
func NewComputePoliciesClientWithBaseURI(baseURI string, subscriptionID string) ComputePoliciesClient {
	return original.NewComputePoliciesClientWithBaseURI(baseURI, subscriptionID)
}
func NewDataLakeStoreAccountsClient(subscriptionID string) DataLakeStoreAccountsClient {
	return original.NewDataLakeStoreAccountsClient(subscriptionID)
}
func NewDataLakeStoreAccountsClientWithBaseURI(baseURI string, subscriptionID string) DataLakeStoreAccountsClient {
	return original.NewDataLakeStoreAccountsClientWithBaseURI(baseURI, subscriptionID)
}
