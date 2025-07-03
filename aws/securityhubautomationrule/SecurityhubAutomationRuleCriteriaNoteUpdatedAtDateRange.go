// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package securityhubautomationrule


type SecurityhubAutomationRuleCriteriaNoteUpdatedAtDateRange struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/securityhub_automation_rule#unit SecurityhubAutomationRule#unit}.
	Unit *string `field:"required" json:"unit" yaml:"unit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/securityhub_automation_rule#value SecurityhubAutomationRule#value}.
	Value *float64 `field:"required" json:"value" yaml:"value"`
}

