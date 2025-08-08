// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkfirewallvpcendpointassociation


type NetworkfirewallVpcEndpointAssociationSubnetMapping struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/networkfirewall_vpc_endpoint_association#subnet_id NetworkfirewallVpcEndpointAssociation#subnet_id}.
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/networkfirewall_vpc_endpoint_association#ip_address_type NetworkfirewallVpcEndpointAssociation#ip_address_type}.
	IpAddressType *string `field:"optional" json:"ipAddressType" yaml:"ipAddressType"`
}

