// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsvpcpeeringconnection


type DataAwsVpcPeeringConnectionFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.68.0/docs/data-sources/vpc_peering_connection#name DataAwsVpcPeeringConnection#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.68.0/docs/data-sources/vpc_peering_connection#values DataAwsVpcPeeringConnection#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

