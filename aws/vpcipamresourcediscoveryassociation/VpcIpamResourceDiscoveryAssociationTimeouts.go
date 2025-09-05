// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpcipamresourcediscoveryassociation


type VpcIpamResourceDiscoveryAssociationTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.12.0/docs/resources/vpc_ipam_resource_discovery_association#create VpcIpamResourceDiscoveryAssociation#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.12.0/docs/resources/vpc_ipam_resource_discovery_association#delete VpcIpamResourceDiscoveryAssociation#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.12.0/docs/resources/vpc_ipam_resource_discovery_association#update VpcIpamResourceDiscoveryAssociation#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

