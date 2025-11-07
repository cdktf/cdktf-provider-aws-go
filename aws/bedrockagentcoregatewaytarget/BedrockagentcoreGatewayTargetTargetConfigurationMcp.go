// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoregatewaytarget


type BedrockagentcoreGatewayTargetTargetConfigurationMcp struct {
	// lambda block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.20.0/docs/resources/bedrockagentcore_gateway_target#lambda BedrockagentcoreGatewayTarget#lambda}
	Lambda interface{} `field:"optional" json:"lambda" yaml:"lambda"`
	// open_api_schema block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.20.0/docs/resources/bedrockagentcore_gateway_target#open_api_schema BedrockagentcoreGatewayTarget#open_api_schema}
	OpenApiSchema interface{} `field:"optional" json:"openApiSchema" yaml:"openApiSchema"`
	// smithy_model block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.20.0/docs/resources/bedrockagentcore_gateway_target#smithy_model BedrockagentcoreGatewayTarget#smithy_model}
	SmithyModel interface{} `field:"optional" json:"smithyModel" yaml:"smithyModel"`
}

