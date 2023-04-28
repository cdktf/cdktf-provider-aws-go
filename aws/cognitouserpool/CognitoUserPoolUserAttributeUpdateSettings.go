package cognitouserpool


type CognitoUserPoolUserAttributeUpdateSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/resources/cognito_user_pool#attributes_require_verification_before_update CognitoUserPool#attributes_require_verification_before_update}.
	AttributesRequireVerificationBeforeUpdate *[]*string `field:"required" json:"attributesRequireVerificationBeforeUpdate" yaml:"attributesRequireVerificationBeforeUpdate"`
}

