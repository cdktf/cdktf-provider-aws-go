package cognitoriskconfiguration


type CognitoRiskConfigurationAccountTakeoverRiskConfiguration struct {
	// actions block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_risk_configuration#actions CognitoRiskConfiguration#actions}
	Actions *CognitoRiskConfigurationAccountTakeoverRiskConfigurationActions `field:"required" json:"actions" yaml:"actions"`
	// notify_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_risk_configuration#notify_configuration CognitoRiskConfiguration#notify_configuration}
	NotifyConfiguration *CognitoRiskConfigurationAccountTakeoverRiskConfigurationNotifyConfiguration `field:"required" json:"notifyConfiguration" yaml:"notifyConfiguration"`
}

