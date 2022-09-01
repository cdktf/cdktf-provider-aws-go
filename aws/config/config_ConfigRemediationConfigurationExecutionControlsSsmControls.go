package config


type ConfigRemediationConfigurationExecutionControlsSsmControls struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_remediation_configuration#concurrent_execution_rate_percentage ConfigRemediationConfiguration#concurrent_execution_rate_percentage}.
	ConcurrentExecutionRatePercentage *float64 `field:"optional" json:"concurrentExecutionRatePercentage" yaml:"concurrentExecutionRatePercentage"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_remediation_configuration#error_percentage ConfigRemediationConfiguration#error_percentage}.
	ErrorPercentage *float64 `field:"optional" json:"errorPercentage" yaml:"errorPercentage"`
}

