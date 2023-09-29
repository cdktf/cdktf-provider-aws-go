// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appflowconnectorprofile


type AppflowConnectorProfileConnectorProfileConfig struct {
	// connector_profile_credentials block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.19.0/docs/resources/appflow_connector_profile#connector_profile_credentials AppflowConnectorProfile#connector_profile_credentials}
	ConnectorProfileCredentials *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentials `field:"required" json:"connectorProfileCredentials" yaml:"connectorProfileCredentials"`
	// connector_profile_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.19.0/docs/resources/appflow_connector_profile#connector_profile_properties AppflowConnectorProfile#connector_profile_properties}
	ConnectorProfileProperties *AppflowConnectorProfileConnectorProfileConfigConnectorProfileProperties `field:"required" json:"connectorProfileProperties" yaml:"connectorProfileProperties"`
}

