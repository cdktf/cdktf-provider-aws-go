// Prebuilt aws Provider for Terraform CDK (cdktf)
package aws


type AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsCustomConnectorApiKey struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appflow_connector_profile#api_key AppflowConnectorProfile#api_key}.
	ApiKey *string `field:"required" json:"apiKey" yaml:"apiKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appflow_connector_profile#api_secret_key AppflowConnectorProfile#api_secret_key}.
	ApiSecretKey *string `field:"optional" json:"apiSecretKey" yaml:"apiSecretKey"`
}

