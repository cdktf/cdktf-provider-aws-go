// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpcencryptioncontrol

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VpcEncryptionControlConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/vpc_encryption_control#mode VpcEncryptionControl#mode}.
	Mode *string `field:"required" json:"mode" yaml:"mode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/vpc_encryption_control#vpc_id VpcEncryptionControl#vpc_id}.
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/vpc_encryption_control#egress_only_internet_gateway_exclusion VpcEncryptionControl#egress_only_internet_gateway_exclusion}.
	EgressOnlyInternetGatewayExclusion *string `field:"optional" json:"egressOnlyInternetGatewayExclusion" yaml:"egressOnlyInternetGatewayExclusion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/vpc_encryption_control#elastic_file_system_exclusion VpcEncryptionControl#elastic_file_system_exclusion}.
	ElasticFileSystemExclusion *string `field:"optional" json:"elasticFileSystemExclusion" yaml:"elasticFileSystemExclusion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/vpc_encryption_control#internet_gateway_exclusion VpcEncryptionControl#internet_gateway_exclusion}.
	InternetGatewayExclusion *string `field:"optional" json:"internetGatewayExclusion" yaml:"internetGatewayExclusion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/vpc_encryption_control#lambda_exclusion VpcEncryptionControl#lambda_exclusion}.
	LambdaExclusion *string `field:"optional" json:"lambdaExclusion" yaml:"lambdaExclusion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/vpc_encryption_control#nat_gateway_exclusion VpcEncryptionControl#nat_gateway_exclusion}.
	NatGatewayExclusion *string `field:"optional" json:"natGatewayExclusion" yaml:"natGatewayExclusion"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/vpc_encryption_control#region VpcEncryptionControl#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/vpc_encryption_control#tags VpcEncryptionControl#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/vpc_encryption_control#timeouts VpcEncryptionControl#timeouts}
	Timeouts *VpcEncryptionControlTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/vpc_encryption_control#virtual_private_gateway_exclusion VpcEncryptionControl#virtual_private_gateway_exclusion}.
	VirtualPrivateGatewayExclusion *string `field:"optional" json:"virtualPrivateGatewayExclusion" yaml:"virtualPrivateGatewayExclusion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/vpc_encryption_control#vpc_lattice_exclusion VpcEncryptionControl#vpc_lattice_exclusion}.
	VpcLatticeExclusion *string `field:"optional" json:"vpcLatticeExclusion" yaml:"vpcLatticeExclusion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/vpc_encryption_control#vpc_peering_exclusion VpcEncryptionControl#vpc_peering_exclusion}.
	VpcPeeringExclusion *string `field:"optional" json:"vpcPeeringExclusion" yaml:"vpcPeeringExclusion"`
}

