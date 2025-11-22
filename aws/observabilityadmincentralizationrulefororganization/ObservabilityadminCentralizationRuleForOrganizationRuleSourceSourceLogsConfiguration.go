// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package observabilityadmincentralizationrulefororganization


type ObservabilityadminCentralizationRuleForOrganizationRuleSourceSourceLogsConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/observabilityadmin_centralization_rule_for_organization#encrypted_log_group_strategy ObservabilityadminCentralizationRuleForOrganization#encrypted_log_group_strategy}.
	EncryptedLogGroupStrategy *string `field:"required" json:"encryptedLogGroupStrategy" yaml:"encryptedLogGroupStrategy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/observabilityadmin_centralization_rule_for_organization#log_group_selection_criteria ObservabilityadminCentralizationRuleForOrganization#log_group_selection_criteria}.
	LogGroupSelectionCriteria *string `field:"required" json:"logGroupSelectionCriteria" yaml:"logGroupSelectionCriteria"`
}

