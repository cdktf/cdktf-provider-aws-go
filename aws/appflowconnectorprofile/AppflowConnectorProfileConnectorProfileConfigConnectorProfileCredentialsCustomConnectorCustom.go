// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appflowconnectorprofile


type AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsCustomConnectorCustom struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/appflow_connector_profile#custom_authentication_type AppflowConnectorProfile#custom_authentication_type}.
	CustomAuthenticationType *string `field:"required" json:"customAuthenticationType" yaml:"customAuthenticationType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/appflow_connector_profile#credentials_map AppflowConnectorProfile#credentials_map}.
	CredentialsMap *map[string]*string `field:"optional" json:"credentialsMap" yaml:"credentialsMap"`
}

