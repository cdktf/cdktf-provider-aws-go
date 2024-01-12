// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigatewayv2integration


type Apigatewayv2IntegrationTlsConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.0/docs/resources/apigatewayv2_integration#server_name_to_verify Apigatewayv2Integration#server_name_to_verify}.
	ServerNameToVerify *string `field:"optional" json:"serverNameToVerify" yaml:"serverNameToVerify"`
}

