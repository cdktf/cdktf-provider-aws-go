// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigatewayv2api


type Apigatewayv2ApiCorsConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.48.0/docs/resources/apigatewayv2_api#allow_credentials Apigatewayv2Api#allow_credentials}.
	AllowCredentials interface{} `field:"optional" json:"allowCredentials" yaml:"allowCredentials"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.48.0/docs/resources/apigatewayv2_api#allow_headers Apigatewayv2Api#allow_headers}.
	AllowHeaders *[]*string `field:"optional" json:"allowHeaders" yaml:"allowHeaders"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.48.0/docs/resources/apigatewayv2_api#allow_methods Apigatewayv2Api#allow_methods}.
	AllowMethods *[]*string `field:"optional" json:"allowMethods" yaml:"allowMethods"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.48.0/docs/resources/apigatewayv2_api#allow_origins Apigatewayv2Api#allow_origins}.
	AllowOrigins *[]*string `field:"optional" json:"allowOrigins" yaml:"allowOrigins"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.48.0/docs/resources/apigatewayv2_api#expose_headers Apigatewayv2Api#expose_headers}.
	ExposeHeaders *[]*string `field:"optional" json:"exposeHeaders" yaml:"exposeHeaders"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.48.0/docs/resources/apigatewayv2_api#max_age Apigatewayv2Api#max_age}.
	MaxAge *float64 `field:"optional" json:"maxAge" yaml:"maxAge"`
}

