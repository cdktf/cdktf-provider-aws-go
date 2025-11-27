// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoregateway


type BedrockagentcoreGatewayProtocolConfigurationMcp struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/bedrockagentcore_gateway#instructions BedrockagentcoreGateway#instructions}.
	Instructions *string `field:"optional" json:"instructions" yaml:"instructions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/bedrockagentcore_gateway#search_type BedrockagentcoreGateway#search_type}.
	SearchType *string `field:"optional" json:"searchType" yaml:"searchType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/bedrockagentcore_gateway#supported_versions BedrockagentcoreGateway#supported_versions}.
	SupportedVersions *[]*string `field:"optional" json:"supportedVersions" yaml:"supportedVersions"`
}

