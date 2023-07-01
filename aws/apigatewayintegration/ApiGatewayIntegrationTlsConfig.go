package apigatewayintegration


type ApiGatewayIntegrationTlsConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/resources/api_gateway_integration#insecure_skip_verification ApiGatewayIntegration#insecure_skip_verification}.
	InsecureSkipVerification interface{} `field:"optional" json:"insecureSkipVerification" yaml:"insecureSkipVerification"`
}

