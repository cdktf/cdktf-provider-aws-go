package apigatewayv2authorizer


type Apigatewayv2AuthorizerJwtConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_authorizer#audience Apigatewayv2Authorizer#audience}.
	Audience *[]*string `field:"optional" json:"audience" yaml:"audience"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apigatewayv2_authorizer#issuer Apigatewayv2Authorizer#issuer}.
	Issuer *string `field:"optional" json:"issuer" yaml:"issuer"`
}

