// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package securityhubautomationrule


type SecurityhubAutomationRuleActionsFindingFieldsUpdateSeverity struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.20.0/docs/resources/securityhub_automation_rule#label SecurityhubAutomationRule#label}.
	Label *string `field:"optional" json:"label" yaml:"label"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.20.0/docs/resources/securityhub_automation_rule#product SecurityhubAutomationRule#product}.
	Product *float64 `field:"optional" json:"product" yaml:"product"`
}

