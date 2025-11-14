// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreoauth2credentialprovider


type BedrockagentcoreOauth2CredentialProviderOauth2ProviderConfig struct {
	// custom_oauth2_provider_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.21.0/docs/resources/bedrockagentcore_oauth2_credential_provider#custom_oauth2_provider_config BedrockagentcoreOauth2CredentialProvider#custom_oauth2_provider_config}
	CustomOauth2ProviderConfig interface{} `field:"optional" json:"customOauth2ProviderConfig" yaml:"customOauth2ProviderConfig"`
	// github_oauth2_provider_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.21.0/docs/resources/bedrockagentcore_oauth2_credential_provider#github_oauth2_provider_config BedrockagentcoreOauth2CredentialProvider#github_oauth2_provider_config}
	GithubOauth2ProviderConfig interface{} `field:"optional" json:"githubOauth2ProviderConfig" yaml:"githubOauth2ProviderConfig"`
	// google_oauth2_provider_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.21.0/docs/resources/bedrockagentcore_oauth2_credential_provider#google_oauth2_provider_config BedrockagentcoreOauth2CredentialProvider#google_oauth2_provider_config}
	GoogleOauth2ProviderConfig interface{} `field:"optional" json:"googleOauth2ProviderConfig" yaml:"googleOauth2ProviderConfig"`
	// microsoft_oauth2_provider_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.21.0/docs/resources/bedrockagentcore_oauth2_credential_provider#microsoft_oauth2_provider_config BedrockagentcoreOauth2CredentialProvider#microsoft_oauth2_provider_config}
	MicrosoftOauth2ProviderConfig interface{} `field:"optional" json:"microsoftOauth2ProviderConfig" yaml:"microsoftOauth2ProviderConfig"`
	// salesforce_oauth2_provider_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.21.0/docs/resources/bedrockagentcore_oauth2_credential_provider#salesforce_oauth2_provider_config BedrockagentcoreOauth2CredentialProvider#salesforce_oauth2_provider_config}
	SalesforceOauth2ProviderConfig interface{} `field:"optional" json:"salesforceOauth2ProviderConfig" yaml:"salesforceOauth2ProviderConfig"`
	// slack_oauth2_provider_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.21.0/docs/resources/bedrockagentcore_oauth2_credential_provider#slack_oauth2_provider_config BedrockagentcoreOauth2CredentialProvider#slack_oauth2_provider_config}
	SlackOauth2ProviderConfig interface{} `field:"optional" json:"slackOauth2ProviderConfig" yaml:"slackOauth2ProviderConfig"`
}

