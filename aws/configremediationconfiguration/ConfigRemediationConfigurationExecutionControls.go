// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package configremediationconfiguration


type ConfigRemediationConfigurationExecutionControls struct {
	// ssm_controls block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/config_remediation_configuration#ssm_controls ConfigRemediationConfiguration#ssm_controls}
	SsmControls *ConfigRemediationConfigurationExecutionControlsSsmControls `field:"optional" json:"ssmControls" yaml:"ssmControls"`
}

