// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoregateway


type BedrockagentcoreGatewayAuthorizerConfigurationCustomJwtAuthorizer struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/bedrockagentcore_gateway#discovery_url BedrockagentcoreGateway#discovery_url}.
	DiscoveryUrl *string `field:"required" json:"discoveryUrl" yaml:"discoveryUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/bedrockagentcore_gateway#allowed_audience BedrockagentcoreGateway#allowed_audience}.
	AllowedAudience *[]*string `field:"optional" json:"allowedAudience" yaml:"allowedAudience"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/bedrockagentcore_gateway#allowed_clients BedrockagentcoreGateway#allowed_clients}.
	AllowedClients *[]*string `field:"optional" json:"allowedClients" yaml:"allowedClients"`
}

