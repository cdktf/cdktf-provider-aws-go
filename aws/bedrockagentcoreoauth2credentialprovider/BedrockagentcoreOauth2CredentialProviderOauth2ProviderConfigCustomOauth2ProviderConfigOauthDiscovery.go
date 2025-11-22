// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreoauth2credentialprovider


type BedrockagentcoreOauth2CredentialProviderOauth2ProviderConfigCustomOauth2ProviderConfigOauthDiscovery struct {
	// authorization_server_metadata block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/bedrockagentcore_oauth2_credential_provider#authorization_server_metadata BedrockagentcoreOauth2CredentialProvider#authorization_server_metadata}
	AuthorizationServerMetadata interface{} `field:"optional" json:"authorizationServerMetadata" yaml:"authorizationServerMetadata"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/bedrockagentcore_oauth2_credential_provider#discovery_url BedrockagentcoreOauth2CredentialProvider#discovery_url}.
	DiscoveryUrl *string `field:"optional" json:"discoveryUrl" yaml:"discoveryUrl"`
}

