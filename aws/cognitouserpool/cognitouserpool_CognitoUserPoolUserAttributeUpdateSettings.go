package cognitouserpool


type CognitoUserPoolUserAttributeUpdateSettings struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool#attributes_require_verification_before_update CognitoUserPool#attributes_require_verification_before_update}.
	AttributesRequireVerificationBeforeUpdate *[]*string `field:"required" json:"attributesRequireVerificationBeforeUpdate" yaml:"attributesRequireVerificationBeforeUpdate"`
}

