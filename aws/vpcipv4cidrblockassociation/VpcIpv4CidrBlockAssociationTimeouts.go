// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpcipv4cidrblockassociation


type VpcIpv4CidrBlockAssociationTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.17.0/docs/resources/vpc_ipv4_cidr_block_association#create VpcIpv4CidrBlockAssociation#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.17.0/docs/resources/vpc_ipv4_cidr_block_association#delete VpcIpv4CidrBlockAssociation#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

