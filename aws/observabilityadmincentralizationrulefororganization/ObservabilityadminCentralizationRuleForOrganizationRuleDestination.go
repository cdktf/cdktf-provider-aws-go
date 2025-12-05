// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package observabilityadmincentralizationrulefororganization


type ObservabilityadminCentralizationRuleForOrganizationRuleDestination struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/observabilityadmin_centralization_rule_for_organization#account ObservabilityadminCentralizationRuleForOrganization#account}.
	Account *string `field:"required" json:"account" yaml:"account"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/observabilityadmin_centralization_rule_for_organization#region ObservabilityadminCentralizationRuleForOrganization#region}.
	Region *string `field:"required" json:"region" yaml:"region"`
	// destination_logs_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/observabilityadmin_centralization_rule_for_organization#destination_logs_configuration ObservabilityadminCentralizationRuleForOrganization#destination_logs_configuration}
	DestinationLogsConfiguration interface{} `field:"optional" json:"destinationLogsConfiguration" yaml:"destinationLogsConfiguration"`
}

