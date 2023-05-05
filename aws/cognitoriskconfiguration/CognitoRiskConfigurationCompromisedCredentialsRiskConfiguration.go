package cognitoriskconfiguration


type CognitoRiskConfigurationCompromisedCredentialsRiskConfiguration struct {
	// actions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/cognito_risk_configuration#actions CognitoRiskConfiguration#actions}
	Actions *CognitoRiskConfigurationCompromisedCredentialsRiskConfigurationActions `field:"required" json:"actions" yaml:"actions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/cognito_risk_configuration#event_filter CognitoRiskConfiguration#event_filter}.
	EventFilter *[]*string `field:"optional" json:"eventFilter" yaml:"eventFilter"`
}

