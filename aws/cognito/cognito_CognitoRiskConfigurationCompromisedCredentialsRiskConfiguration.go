package cognito


type CognitoRiskConfigurationCompromisedCredentialsRiskConfiguration struct {
	// actions block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_risk_configuration#actions CognitoRiskConfiguration#actions}
	Actions *CognitoRiskConfigurationCompromisedCredentialsRiskConfigurationActions `field:"required" json:"actions" yaml:"actions"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_risk_configuration#event_filter CognitoRiskConfiguration#event_filter}.
	EventFilter *[]*string `field:"optional" json:"eventFilter" yaml:"eventFilter"`
}

