// Prebuilt aws Provider for Terraform CDK (cdktf)
package aws


type AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsCustomConnectorCustom struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appflow_connector_profile#custom_authentication_type AppflowConnectorProfile#custom_authentication_type}.
	CustomAuthenticationType *string `field:"required" json:"customAuthenticationType" yaml:"customAuthenticationType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appflow_connector_profile#credentials_map AppflowConnectorProfile#credentials_map}.
	CredentialsMap *map[string]*string `field:"optional" json:"credentialsMap" yaml:"credentialsMap"`
}

