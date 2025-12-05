// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package observabilityadmincentralizationrulefororganization


type ObservabilityadminCentralizationRuleForOrganizationRuleSource struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/observabilityadmin_centralization_rule_for_organization#regions ObservabilityadminCentralizationRuleForOrganization#regions}.
	Regions *[]*string `field:"required" json:"regions" yaml:"regions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/observabilityadmin_centralization_rule_for_organization#scope ObservabilityadminCentralizationRuleForOrganization#scope}.
	Scope *string `field:"required" json:"scope" yaml:"scope"`
	// source_logs_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/observabilityadmin_centralization_rule_for_organization#source_logs_configuration ObservabilityadminCentralizationRuleForOrganization#source_logs_configuration}
	SourceLogsConfiguration interface{} `field:"optional" json:"sourceLogsConfiguration" yaml:"sourceLogsConfiguration"`
}

