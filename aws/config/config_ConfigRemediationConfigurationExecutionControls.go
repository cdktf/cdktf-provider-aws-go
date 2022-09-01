package config


type ConfigRemediationConfigurationExecutionControls struct {
	// ssm_controls block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_remediation_configuration#ssm_controls ConfigRemediationConfiguration#ssm_controls}
	SsmControls *ConfigRemediationConfigurationExecutionControlsSsmControls `field:"optional" json:"ssmControls" yaml:"ssmControls"`
}

