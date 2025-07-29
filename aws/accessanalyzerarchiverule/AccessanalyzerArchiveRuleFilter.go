// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accessanalyzerarchiverule


type AccessanalyzerArchiveRuleFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.6.0/docs/resources/accessanalyzer_archive_rule#criteria AccessanalyzerArchiveRule#criteria}.
	Criteria *string `field:"required" json:"criteria" yaml:"criteria"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.6.0/docs/resources/accessanalyzer_archive_rule#contains AccessanalyzerArchiveRule#contains}.
	Contains *[]*string `field:"optional" json:"contains" yaml:"contains"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.6.0/docs/resources/accessanalyzer_archive_rule#eq AccessanalyzerArchiveRule#eq}.
	Eq *[]*string `field:"optional" json:"eq" yaml:"eq"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.6.0/docs/resources/accessanalyzer_archive_rule#exists AccessanalyzerArchiveRule#exists}.
	Exists *string `field:"optional" json:"exists" yaml:"exists"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.6.0/docs/resources/accessanalyzer_archive_rule#neq AccessanalyzerArchiveRule#neq}.
	Neq *[]*string `field:"optional" json:"neq" yaml:"neq"`
}

