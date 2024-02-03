// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigatewayv2authorizer


type Apigatewayv2AuthorizerJwtConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/resources/apigatewayv2_authorizer#audience Apigatewayv2Authorizer#audience}.
	Audience *[]*string `field:"optional" json:"audience" yaml:"audience"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/resources/apigatewayv2_authorizer#issuer Apigatewayv2Authorizer#issuer}.
	Issuer *string `field:"optional" json:"issuer" yaml:"issuer"`
}

