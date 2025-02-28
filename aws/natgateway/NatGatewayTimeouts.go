// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package natgateway


type NatGatewayTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/nat_gateway#create NatGateway#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/nat_gateway#delete NatGateway#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/nat_gateway#update NatGateway#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

