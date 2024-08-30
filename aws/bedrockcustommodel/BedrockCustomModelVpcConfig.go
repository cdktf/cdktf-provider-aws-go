// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockcustommodel


type BedrockCustomModelVpcConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.65.0/docs/resources/bedrock_custom_model#security_group_ids BedrockCustomModel#security_group_ids}.
	SecurityGroupIds *[]*string `field:"required" json:"securityGroupIds" yaml:"securityGroupIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.65.0/docs/resources/bedrock_custom_model#subnet_ids BedrockCustomModel#subnet_ids}.
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
}

