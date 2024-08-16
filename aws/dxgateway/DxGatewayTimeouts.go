// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dxgateway


type DxGatewayTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/dx_gateway#create DxGateway#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/dx_gateway#delete DxGateway#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

