// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accessanalyzeranalyzer


type AccessanalyzerAnalyzerConfigurationInternalAccessAnalysisRuleInclusion struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/accessanalyzer_analyzer#account_ids AccessanalyzerAnalyzer#account_ids}.
	AccountIds *[]*string `field:"optional" json:"accountIds" yaml:"accountIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/accessanalyzer_analyzer#resource_arns AccessanalyzerAnalyzer#resource_arns}.
	ResourceArns *[]*string `field:"optional" json:"resourceArns" yaml:"resourceArns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/accessanalyzer_analyzer#resource_types AccessanalyzerAnalyzer#resource_types}.
	ResourceTypes *[]*string `field:"optional" json:"resourceTypes" yaml:"resourceTypes"`
}

