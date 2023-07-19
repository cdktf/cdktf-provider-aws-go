package cognitouserpool


type CognitoUserPoolAccountRecoverySettingRecoveryMechanism struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.8.0/docs/resources/cognito_user_pool#name CognitoUserPool#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.8.0/docs/resources/cognito_user_pool#priority CognitoUserPool#priority}.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
}

