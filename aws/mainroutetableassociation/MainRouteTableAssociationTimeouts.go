// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mainroutetableassociation


type MainRouteTableAssociationTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.2/docs/resources/main_route_table_association#create MainRouteTableAssociation#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.2/docs/resources/main_route_table_association#delete MainRouteTableAssociation#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.2/docs/resources/main_route_table_association#update MainRouteTableAssociation#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

