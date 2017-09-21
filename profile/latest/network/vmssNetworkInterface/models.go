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

package vmssnetworkinterface

import original "github.com/Azure/azure-sdk-for-go/service/network/management/2017-03-01/vmssNetworkInterface"

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ManagementClient = original.ManagementClient
type IPAllocationMethod = original.IPAllocationMethod

const (
	Dynamic	IPAllocationMethod	= original.Dynamic
	Static	IPAllocationMethod	= original.Static
)

type IPVersion = original.IPVersion

const (
	IPv4	IPVersion	= original.IPv4
	IPv6	IPVersion	= original.IPv6
)

type RouteNextHopType = original.RouteNextHopType

const (
	Internet		RouteNextHopType	= original.Internet
	None			RouteNextHopType	= original.None
	VirtualAppliance	RouteNextHopType	= original.VirtualAppliance
	VirtualNetworkGateway	RouteNextHopType	= original.VirtualNetworkGateway
	VnetLocal		RouteNextHopType	= original.VnetLocal
)

type SecurityRuleAccess = original.SecurityRuleAccess

const (
	Allow	SecurityRuleAccess	= original.Allow
	Deny	SecurityRuleAccess	= original.Deny
)

type SecurityRuleDirection = original.SecurityRuleDirection

const (
	Inbound		SecurityRuleDirection	= original.Inbound
	Outbound	SecurityRuleDirection	= original.Outbound
)

type SecurityRuleProtocol = original.SecurityRuleProtocol

const (
	Asterisk	SecurityRuleProtocol	= original.Asterisk
	TCP		SecurityRuleProtocol	= original.TCP
	UDP		SecurityRuleProtocol	= original.UDP
)

type TransportProtocol = original.TransportProtocol

const (
	TransportProtocolTCP	TransportProtocol	= original.TransportProtocolTCP
	TransportProtocolUDP	TransportProtocol	= original.TransportProtocolUDP
)

type ApplicationGatewayBackendAddress = original.ApplicationGatewayBackendAddress
type ApplicationGatewayBackendAddressPool = original.ApplicationGatewayBackendAddressPool
type ApplicationGatewayBackendAddressPoolPropertiesFormat = original.ApplicationGatewayBackendAddressPoolPropertiesFormat
type BackendAddressPool = original.BackendAddressPool
type BackendAddressPoolPropertiesFormat = original.BackendAddressPoolPropertiesFormat
type InboundNatRule = original.InboundNatRule
type InboundNatRulePropertiesFormat = original.InboundNatRulePropertiesFormat
type IPConfiguration = original.IPConfiguration
type IPConfigurationPropertiesFormat = original.IPConfigurationPropertiesFormat
type NetworkInterface = original.NetworkInterface
type NetworkInterfaceDNSSettings = original.NetworkInterfaceDNSSettings
type NetworkInterfaceIPConfiguration = original.NetworkInterfaceIPConfiguration
type NetworkInterfaceIPConfigurationPropertiesFormat = original.NetworkInterfaceIPConfigurationPropertiesFormat
type NetworkInterfaceListResult = original.NetworkInterfaceListResult
type NetworkInterfacePropertiesFormat = original.NetworkInterfacePropertiesFormat
type NetworkSecurityGroup = original.NetworkSecurityGroup
type NetworkSecurityGroupPropertiesFormat = original.NetworkSecurityGroupPropertiesFormat
type PublicIPAddress = original.PublicIPAddress
type PublicIPAddressDNSSettings = original.PublicIPAddressDNSSettings
type PublicIPAddressPropertiesFormat = original.PublicIPAddressPropertiesFormat
type Resource = original.Resource
type ResourceNavigationLink = original.ResourceNavigationLink
type ResourceNavigationLinkFormat = original.ResourceNavigationLinkFormat
type Route = original.Route
type RoutePropertiesFormat = original.RoutePropertiesFormat
type RouteTable = original.RouteTable
type RouteTablePropertiesFormat = original.RouteTablePropertiesFormat
type SecurityRule = original.SecurityRule
type SecurityRulePropertiesFormat = original.SecurityRulePropertiesFormat
type Subnet = original.Subnet
type SubnetPropertiesFormat = original.SubnetPropertiesFormat
type SubResource = original.SubResource
type NetworkInterfacesClient = original.NetworkInterfacesClient

func New(subscriptionID string) ManagementClient {
	return original.New(subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) ManagementClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func NewNetworkInterfacesClient(subscriptionID string) NetworkInterfacesClient {
	return original.NewNetworkInterfacesClient(subscriptionID)
}
func NewNetworkInterfacesClientWithBaseURI(baseURI string, subscriptionID string) NetworkInterfacesClient {
	return original.NewNetworkInterfacesClientWithBaseURI(baseURI, subscriptionID)
}
func UserAgent() string {
	return original.UserAgent()
}
func Version() string {
	return original.Version()
}
