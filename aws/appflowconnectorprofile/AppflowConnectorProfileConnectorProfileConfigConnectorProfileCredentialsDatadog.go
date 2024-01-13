// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appflowconnectorprofile


type AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsDatadog struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.1/docs/resources/appflow_connector_profile#api_key AppflowConnectorProfile#api_key}.
	ApiKey *string `field:"required" json:"apiKey" yaml:"apiKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.1/docs/resources/appflow_connector_profile#application_key AppflowConnectorProfile#application_key}.
	ApplicationKey *string `field:"required" json:"applicationKey" yaml:"applicationKey"`
}

