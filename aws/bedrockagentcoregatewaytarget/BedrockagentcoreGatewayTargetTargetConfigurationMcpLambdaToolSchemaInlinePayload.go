// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoregatewaytarget


type BedrockagentcoreGatewayTargetTargetConfigurationMcpLambdaToolSchemaInlinePayload struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/bedrockagentcore_gateway_target#description BedrockagentcoreGatewayTarget#description}.
	Description *string `field:"required" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/bedrockagentcore_gateway_target#name BedrockagentcoreGatewayTarget#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// input_schema block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/bedrockagentcore_gateway_target#input_schema BedrockagentcoreGatewayTarget#input_schema}
	InputSchema interface{} `field:"optional" json:"inputSchema" yaml:"inputSchema"`
	// output_schema block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/bedrockagentcore_gateway_target#output_schema BedrockagentcoreGatewayTarget#output_schema}
	OutputSchema interface{} `field:"optional" json:"outputSchema" yaml:"outputSchema"`
}

