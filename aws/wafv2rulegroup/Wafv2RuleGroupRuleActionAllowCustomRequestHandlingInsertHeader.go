// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafv2rulegroup


type Wafv2RuleGroupRuleActionAllowCustomRequestHandlingInsertHeader struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.34.0/docs/resources/wafv2_rule_group#name Wafv2RuleGroup#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.34.0/docs/resources/wafv2_rule_group#value Wafv2RuleGroup#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

