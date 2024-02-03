// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsec2localgatewayvirtualinterfacegroups


type DataAwsEc2LocalGatewayVirtualInterfaceGroupsFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/data-sources/ec2_local_gateway_virtual_interface_groups#name DataAwsEc2LocalGatewayVirtualInterfaceGroups#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/data-sources/ec2_local_gateway_virtual_interface_groups#values DataAwsEc2LocalGatewayVirtualInterfaceGroups#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

