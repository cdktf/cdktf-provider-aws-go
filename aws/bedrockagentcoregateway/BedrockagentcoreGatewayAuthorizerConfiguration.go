// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoregateway


type BedrockagentcoreGatewayAuthorizerConfiguration struct {
	// custom_jwt_authorizer block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/bedrockagentcore_gateway#custom_jwt_authorizer BedrockagentcoreGateway#custom_jwt_authorizer}
	CustomJwtAuthorizer interface{} `field:"optional" json:"customJwtAuthorizer" yaml:"customJwtAuthorizer"`
}

