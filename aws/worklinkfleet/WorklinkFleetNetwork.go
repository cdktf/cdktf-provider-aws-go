// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package worklinkfleet


type WorklinkFleetNetwork struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.1/docs/resources/worklink_fleet#security_group_ids WorklinkFleet#security_group_ids}.
	SecurityGroupIds *[]*string `field:"required" json:"securityGroupIds" yaml:"securityGroupIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.1/docs/resources/worklink_fleet#subnet_ids WorklinkFleet#subnet_ids}.
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.1/docs/resources/worklink_fleet#vpc_id WorklinkFleet#vpc_id}.
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
}

