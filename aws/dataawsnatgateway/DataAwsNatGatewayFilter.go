// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsnatgateway


type DataAwsNatGatewayFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.1/docs/data-sources/nat_gateway#name DataAwsNatGateway#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.1/docs/data-sources/nat_gateway#values DataAwsNatGateway#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

