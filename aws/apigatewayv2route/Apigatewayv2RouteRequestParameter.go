// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigatewayv2route


type Apigatewayv2RouteRequestParameter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.72.1/docs/resources/apigatewayv2_route#request_parameter_key Apigatewayv2Route#request_parameter_key}.
	RequestParameterKey *string `field:"required" json:"requestParameterKey" yaml:"requestParameterKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.72.1/docs/resources/apigatewayv2_route#required Apigatewayv2Route#required}.
	Required interface{} `field:"required" json:"required" yaml:"required"`
}

