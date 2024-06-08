// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsinternetgateway


type DataAwsInternetGatewayFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.53.0/docs/data-sources/internet_gateway#name DataAwsInternetGateway#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.53.0/docs/data-sources/internet_gateway#values DataAwsInternetGateway#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

