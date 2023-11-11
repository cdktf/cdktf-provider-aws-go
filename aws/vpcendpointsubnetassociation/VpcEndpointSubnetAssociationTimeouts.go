// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpcendpointsubnetassociation


type VpcEndpointSubnetAssociationTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.25.0/docs/resources/vpc_endpoint_subnet_association#create VpcEndpointSubnetAssociation#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.25.0/docs/resources/vpc_endpoint_subnet_association#delete VpcEndpointSubnetAssociation#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

