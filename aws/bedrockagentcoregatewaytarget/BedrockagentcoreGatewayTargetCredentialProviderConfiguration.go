// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoregatewaytarget


type BedrockagentcoreGatewayTargetCredentialProviderConfiguration struct {
	// api_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/bedrockagentcore_gateway_target#api_key BedrockagentcoreGatewayTarget#api_key}
	ApiKey interface{} `field:"optional" json:"apiKey" yaml:"apiKey"`
	// gateway_iam_role block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/bedrockagentcore_gateway_target#gateway_iam_role BedrockagentcoreGatewayTarget#gateway_iam_role}
	GatewayIamRole interface{} `field:"optional" json:"gatewayIamRole" yaml:"gatewayIamRole"`
	// oauth block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/bedrockagentcore_gateway_target#oauth BedrockagentcoreGatewayTarget#oauth}
	Oauth interface{} `field:"optional" json:"oauth" yaml:"oauth"`
}

