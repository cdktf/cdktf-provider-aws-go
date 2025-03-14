// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package osispipeline


type OsisPipelineVpcOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.91.0/docs/resources/osis_pipeline#subnet_ids OsisPipeline#subnet_ids}.
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.91.0/docs/resources/osis_pipeline#security_group_ids OsisPipeline#security_group_ids}.
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.91.0/docs/resources/osis_pipeline#vpc_endpoint_management OsisPipeline#vpc_endpoint_management}.
	VpcEndpointManagement *string `field:"optional" json:"vpcEndpointManagement" yaml:"vpcEndpointManagement"`
}

