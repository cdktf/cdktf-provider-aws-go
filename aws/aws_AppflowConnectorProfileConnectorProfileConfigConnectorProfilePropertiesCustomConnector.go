// Prebuilt aws Provider for Terraform CDK (cdktf)
package aws


type AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesCustomConnector struct {
	// oauth2_properties block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appflow_connector_profile#oauth2_properties AppflowConnectorProfile#oauth2_properties}
	Oauth2Properties *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesCustomConnectorOauth2Properties `field:"optional" json:"oauth2Properties" yaml:"oauth2Properties"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appflow_connector_profile#profile_properties AppflowConnectorProfile#profile_properties}.
	ProfileProperties *map[string]*string `field:"optional" json:"profileProperties" yaml:"profileProperties"`
}

