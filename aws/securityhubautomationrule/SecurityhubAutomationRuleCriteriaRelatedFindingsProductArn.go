// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package securityhubautomationrule


type SecurityhubAutomationRuleCriteriaRelatedFindingsProductArn struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/securityhub_automation_rule#comparison SecurityhubAutomationRule#comparison}.
	Comparison *string `field:"required" json:"comparison" yaml:"comparison"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/securityhub_automation_rule#value SecurityhubAutomationRule#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

