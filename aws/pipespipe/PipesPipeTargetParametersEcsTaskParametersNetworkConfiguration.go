// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pipespipe


type PipesPipeTargetParametersEcsTaskParametersNetworkConfiguration struct {
	// aws_vpc_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.20.0/docs/resources/pipes_pipe#aws_vpc_configuration PipesPipe#aws_vpc_configuration}
	AwsVpcConfiguration *PipesPipeTargetParametersEcsTaskParametersNetworkConfigurationAwsVpcConfiguration `field:"optional" json:"awsVpcConfiguration" yaml:"awsVpcConfiguration"`
}

