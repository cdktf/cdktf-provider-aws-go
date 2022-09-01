// Prebuilt aws Provider for Terraform CDK (cdktf)
package aws


type AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsCustomConnector struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appflow_connector_profile#authentication_type AppflowConnectorProfile#authentication_type}.
	AuthenticationType *string `field:"required" json:"authenticationType" yaml:"authenticationType"`
	// api_key block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appflow_connector_profile#api_key AppflowConnectorProfile#api_key}
	ApiKey *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsCustomConnectorApiKey `field:"optional" json:"apiKey" yaml:"apiKey"`
	// basic block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appflow_connector_profile#basic AppflowConnectorProfile#basic}
	Basic *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsCustomConnectorBasic `field:"optional" json:"basic" yaml:"basic"`
	// custom block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appflow_connector_profile#custom AppflowConnectorProfile#custom}
	Custom *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsCustomConnectorCustom `field:"optional" json:"custom" yaml:"custom"`
	// oauth2 block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appflow_connector_profile#oauth2 AppflowConnectorProfile#oauth2}
	Oauth2 *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsCustomConnectorOauth2 `field:"optional" json:"oauth2" yaml:"oauth2"`
}

