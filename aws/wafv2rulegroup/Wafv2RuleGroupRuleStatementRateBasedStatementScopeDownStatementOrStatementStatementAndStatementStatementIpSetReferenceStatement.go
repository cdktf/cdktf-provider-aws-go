// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafv2rulegroup


type Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementIpSetReferenceStatement struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/wafv2_rule_group#arn Wafv2RuleGroup#arn}.
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// ip_set_forwarded_ip_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/wafv2_rule_group#ip_set_forwarded_ip_config Wafv2RuleGroup#ip_set_forwarded_ip_config}
	IpSetForwardedIpConfig *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementIpSetReferenceStatementIpSetForwardedIpConfig `field:"optional" json:"ipSetForwardedIpConfig" yaml:"ipSetForwardedIpConfig"`
}

