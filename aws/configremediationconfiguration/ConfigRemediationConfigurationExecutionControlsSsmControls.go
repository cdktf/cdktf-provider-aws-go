// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package configremediationconfiguration


type ConfigRemediationConfigurationExecutionControlsSsmControls struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.11.0/docs/resources/config_remediation_configuration#concurrent_execution_rate_percentage ConfigRemediationConfiguration#concurrent_execution_rate_percentage}.
	ConcurrentExecutionRatePercentage *float64 `field:"optional" json:"concurrentExecutionRatePercentage" yaml:"concurrentExecutionRatePercentage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.11.0/docs/resources/config_remediation_configuration#error_percentage ConfigRemediationConfiguration#error_percentage}.
	ErrorPercentage *float64 `field:"optional" json:"errorPercentage" yaml:"errorPercentage"`
}

