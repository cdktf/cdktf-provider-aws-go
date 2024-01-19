// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appflowconnectorprofile


type AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSapoData struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/appflow_connector_profile#application_host_url AppflowConnectorProfile#application_host_url}.
	ApplicationHostUrl *string `field:"required" json:"applicationHostUrl" yaml:"applicationHostUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/appflow_connector_profile#application_service_path AppflowConnectorProfile#application_service_path}.
	ApplicationServicePath *string `field:"required" json:"applicationServicePath" yaml:"applicationServicePath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/appflow_connector_profile#client_number AppflowConnectorProfile#client_number}.
	ClientNumber *string `field:"required" json:"clientNumber" yaml:"clientNumber"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/appflow_connector_profile#port_number AppflowConnectorProfile#port_number}.
	PortNumber *float64 `field:"required" json:"portNumber" yaml:"portNumber"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/appflow_connector_profile#logon_language AppflowConnectorProfile#logon_language}.
	LogonLanguage *string `field:"optional" json:"logonLanguage" yaml:"logonLanguage"`
	// oauth_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/appflow_connector_profile#oauth_properties AppflowConnectorProfile#oauth_properties}
	OauthProperties *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSapoDataOauthProperties `field:"optional" json:"oauthProperties" yaml:"oauthProperties"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/appflow_connector_profile#private_link_service_name AppflowConnectorProfile#private_link_service_name}.
	PrivateLinkServiceName *string `field:"optional" json:"privateLinkServiceName" yaml:"privateLinkServiceName"`
}

