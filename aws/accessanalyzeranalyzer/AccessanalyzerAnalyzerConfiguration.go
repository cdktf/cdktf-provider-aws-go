// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accessanalyzeranalyzer


type AccessanalyzerAnalyzerConfiguration struct {
	// internal_access block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/accessanalyzer_analyzer#internal_access AccessanalyzerAnalyzer#internal_access}
	InternalAccess *AccessanalyzerAnalyzerConfigurationInternalAccess `field:"optional" json:"internalAccess" yaml:"internalAccess"`
	// unused_access block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/accessanalyzer_analyzer#unused_access AccessanalyzerAnalyzer#unused_access}
	UnusedAccess *AccessanalyzerAnalyzerConfigurationUnusedAccess `field:"optional" json:"unusedAccess" yaml:"unusedAccess"`
}

