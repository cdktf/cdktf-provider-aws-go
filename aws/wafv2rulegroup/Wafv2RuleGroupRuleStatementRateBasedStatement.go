// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafv2rulegroup


type Wafv2RuleGroupRuleStatementRateBasedStatement struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.30.0/docs/resources/wafv2_rule_group#limit Wafv2RuleGroup#limit}.
	Limit *float64 `field:"required" json:"limit" yaml:"limit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.30.0/docs/resources/wafv2_rule_group#aggregate_key_type Wafv2RuleGroup#aggregate_key_type}.
	AggregateKeyType *string `field:"optional" json:"aggregateKeyType" yaml:"aggregateKeyType"`
	// custom_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.30.0/docs/resources/wafv2_rule_group#custom_key Wafv2RuleGroup#custom_key}
	CustomKey interface{} `field:"optional" json:"customKey" yaml:"customKey"`
	// forwarded_ip_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.30.0/docs/resources/wafv2_rule_group#forwarded_ip_config Wafv2RuleGroup#forwarded_ip_config}
	ForwardedIpConfig *Wafv2RuleGroupRuleStatementRateBasedStatementForwardedIpConfig `field:"optional" json:"forwardedIpConfig" yaml:"forwardedIpConfig"`
	// scope_down_statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.30.0/docs/resources/wafv2_rule_group#scope_down_statement Wafv2RuleGroup#scope_down_statement}
	ScopeDownStatement *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatement `field:"optional" json:"scopeDownStatement" yaml:"scopeDownStatement"`
}

