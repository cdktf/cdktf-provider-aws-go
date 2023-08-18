package configremediationconfiguration


type ConfigRemediationConfigurationExecutionControls struct {
	// ssm_controls block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/config_remediation_configuration#ssm_controls ConfigRemediationConfiguration#ssm_controls}
	SsmControls *ConfigRemediationConfigurationExecutionControlsSsmControls `field:"optional" json:"ssmControls" yaml:"ssmControls"`
}

