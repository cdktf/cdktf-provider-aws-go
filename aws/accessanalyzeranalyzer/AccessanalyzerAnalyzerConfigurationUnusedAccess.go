// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accessanalyzeranalyzer


type AccessanalyzerAnalyzerConfigurationUnusedAccess struct {
	// analysis_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/accessanalyzer_analyzer#analysis_rule AccessanalyzerAnalyzer#analysis_rule}
	AnalysisRule *AccessanalyzerAnalyzerConfigurationUnusedAccessAnalysisRule `field:"optional" json:"analysisRule" yaml:"analysisRule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/accessanalyzer_analyzer#unused_access_age AccessanalyzerAnalyzer#unused_access_age}.
	UnusedAccessAge *float64 `field:"optional" json:"unusedAccessAge" yaml:"unusedAccessAge"`
}

