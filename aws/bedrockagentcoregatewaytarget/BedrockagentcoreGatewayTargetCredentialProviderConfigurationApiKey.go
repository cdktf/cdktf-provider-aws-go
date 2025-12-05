// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoregatewaytarget


type BedrockagentcoreGatewayTargetCredentialProviderConfigurationApiKey struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/bedrockagentcore_gateway_target#provider_arn BedrockagentcoreGatewayTarget#provider_arn}.
	ProviderArn *string `field:"required" json:"providerArn" yaml:"providerArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/bedrockagentcore_gateway_target#credential_location BedrockagentcoreGatewayTarget#credential_location}.
	CredentialLocation *string `field:"optional" json:"credentialLocation" yaml:"credentialLocation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/bedrockagentcore_gateway_target#credential_parameter_name BedrockagentcoreGatewayTarget#credential_parameter_name}.
	CredentialParameterName *string `field:"optional" json:"credentialParameterName" yaml:"credentialParameterName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/bedrockagentcore_gateway_target#credential_prefix BedrockagentcoreGatewayTarget#credential_prefix}.
	CredentialPrefix *string `field:"optional" json:"credentialPrefix" yaml:"credentialPrefix"`
}

