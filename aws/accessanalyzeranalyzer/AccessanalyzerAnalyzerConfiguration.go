// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accessanalyzeranalyzer


type AccessanalyzerAnalyzerConfiguration struct {
	// unused_access block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.69.0/docs/resources/accessanalyzer_analyzer#unused_access AccessanalyzerAnalyzer#unused_access}
	UnusedAccess *AccessanalyzerAnalyzerConfigurationUnusedAccess `field:"optional" json:"unusedAccess" yaml:"unusedAccess"`
}

