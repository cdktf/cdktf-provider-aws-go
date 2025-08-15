// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appflowconnectorprofile


type AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesCustomConnector struct {
	// oauth2_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/resources/appflow_connector_profile#oauth2_properties AppflowConnectorProfile#oauth2_properties}
	Oauth2Properties *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesCustomConnectorOauth2Properties `field:"optional" json:"oauth2Properties" yaml:"oauth2Properties"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/resources/appflow_connector_profile#profile_properties AppflowConnectorProfile#profile_properties}.
	ProfileProperties *map[string]*string `field:"optional" json:"profileProperties" yaml:"profileProperties"`
}

