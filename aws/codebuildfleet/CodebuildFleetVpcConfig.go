// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package codebuildfleet


type CodebuildFleetVpcConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/codebuild_fleet#security_group_ids CodebuildFleet#security_group_ids}.
	SecurityGroupIds *[]*string `field:"required" json:"securityGroupIds" yaml:"securityGroupIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/codebuild_fleet#subnets CodebuildFleet#subnets}.
	Subnets *[]*string `field:"required" json:"subnets" yaml:"subnets"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/codebuild_fleet#vpc_id CodebuildFleet#vpc_id}.
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
}

