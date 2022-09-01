// Prebuilt aws Provider for Terraform CDK (cdktf)
package aws


type AppflowConnectorProfileConnectorProfileConfig struct {
	// connector_profile_credentials block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appflow_connector_profile#connector_profile_credentials AppflowConnectorProfile#connector_profile_credentials}
	ConnectorProfileCredentials *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentials `field:"required" json:"connectorProfileCredentials" yaml:"connectorProfileCredentials"`
	// connector_profile_properties block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appflow_connector_profile#connector_profile_properties AppflowConnectorProfile#connector_profile_properties}
	ConnectorProfileProperties *AppflowConnectorProfileConnectorProfileConfigConnectorProfileProperties `field:"required" json:"connectorProfileProperties" yaml:"connectorProfileProperties"`
}

