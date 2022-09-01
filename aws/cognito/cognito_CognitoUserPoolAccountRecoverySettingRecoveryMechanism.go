package cognito


type CognitoUserPoolAccountRecoverySettingRecoveryMechanism struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool#name CognitoUserPool#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool#priority CognitoUserPool#priority}.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
}

