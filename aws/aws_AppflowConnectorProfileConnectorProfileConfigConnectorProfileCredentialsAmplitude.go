// Prebuilt aws Provider for Terraform CDK (cdktf)
package aws


type AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsAmplitude struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appflow_connector_profile#api_key AppflowConnectorProfile#api_key}.
	ApiKey *string `field:"required" json:"apiKey" yaml:"apiKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appflow_connector_profile#secret_key AppflowConnectorProfile#secret_key}.
	SecretKey *string `field:"required" json:"secretKey" yaml:"secretKey"`
}

