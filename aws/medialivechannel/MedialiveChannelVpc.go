// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package medialivechannel


type MedialiveChannelVpc struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/medialive_channel#public_address_allocation_ids MedialiveChannel#public_address_allocation_ids}.
	PublicAddressAllocationIds *[]*string `field:"required" json:"publicAddressAllocationIds" yaml:"publicAddressAllocationIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/medialive_channel#subnet_ids MedialiveChannel#subnet_ids}.
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/medialive_channel#security_group_ids MedialiveChannel#security_group_ids}.
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
}

