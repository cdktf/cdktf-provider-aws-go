package cognitouserpool


type CognitoUserPoolAccountRecoverySetting struct {
	// recovery_mechanism block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool#recovery_mechanism CognitoUserPool#recovery_mechanism}
	RecoveryMechanism interface{} `field:"required" json:"recoveryMechanism" yaml:"recoveryMechanism"`
}

