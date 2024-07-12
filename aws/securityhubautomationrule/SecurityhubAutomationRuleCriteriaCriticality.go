// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package securityhubautomationrule


type SecurityhubAutomationRuleCriteriaCriticality struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.58.0/docs/resources/securityhub_automation_rule#eq SecurityhubAutomationRule#eq}.
	Eq *float64 `field:"optional" json:"eq" yaml:"eq"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.58.0/docs/resources/securityhub_automation_rule#gt SecurityhubAutomationRule#gt}.
	Gt *float64 `field:"optional" json:"gt" yaml:"gt"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.58.0/docs/resources/securityhub_automation_rule#gte SecurityhubAutomationRule#gte}.
	Gte *float64 `field:"optional" json:"gte" yaml:"gte"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.58.0/docs/resources/securityhub_automation_rule#lt SecurityhubAutomationRule#lt}.
	Lt *float64 `field:"optional" json:"lt" yaml:"lt"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.58.0/docs/resources/securityhub_automation_rule#lte SecurityhubAutomationRule#lte}.
	Lte *float64 `field:"optional" json:"lte" yaml:"lte"`
}

