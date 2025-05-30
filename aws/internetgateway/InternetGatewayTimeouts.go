// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package internetgateway


type InternetGatewayTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/resources/internet_gateway#create InternetGateway#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/resources/internet_gateway#delete InternetGateway#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/resources/internet_gateway#update InternetGateway#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

