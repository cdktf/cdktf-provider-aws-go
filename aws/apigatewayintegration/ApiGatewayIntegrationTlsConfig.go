package apigatewayintegration


type ApiGatewayIntegrationTlsConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/api_gateway_integration#insecure_skip_verification ApiGatewayIntegration#insecure_skip_verification}.
	InsecureSkipVerification interface{} `field:"optional" json:"insecureSkipVerification" yaml:"insecureSkipVerification"`
}

