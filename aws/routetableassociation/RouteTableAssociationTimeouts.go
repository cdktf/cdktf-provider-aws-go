// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package routetableassociation


type RouteTableAssociationTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.1/docs/resources/route_table_association#create RouteTableAssociation#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.1/docs/resources/route_table_association#delete RouteTableAssociation#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.1/docs/resources/route_table_association#update RouteTableAssociation#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

