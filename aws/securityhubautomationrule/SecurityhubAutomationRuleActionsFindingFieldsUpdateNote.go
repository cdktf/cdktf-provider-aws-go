// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package securityhubautomationrule


type SecurityhubAutomationRuleActionsFindingFieldsUpdateNote struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/securityhub_automation_rule#text SecurityhubAutomationRule#text}.
	Text *string `field:"required" json:"text" yaml:"text"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/securityhub_automation_rule#updated_by SecurityhubAutomationRule#updated_by}.
	UpdatedBy *string `field:"required" json:"updatedBy" yaml:"updatedBy"`
}

