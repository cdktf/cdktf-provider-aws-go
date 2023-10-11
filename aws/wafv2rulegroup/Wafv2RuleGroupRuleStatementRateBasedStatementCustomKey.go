// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafv2rulegroup


type Wafv2RuleGroupRuleStatementRateBasedStatementCustomKey struct {
	// cookie block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.20.1/docs/resources/wafv2_rule_group#cookie Wafv2RuleGroup#cookie}
	Cookie *Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyCookie `field:"optional" json:"cookie" yaml:"cookie"`
	// forwarded_ip block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.20.1/docs/resources/wafv2_rule_group#forwarded_ip Wafv2RuleGroup#forwarded_ip}
	ForwardedIp *Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyForwardedIp `field:"optional" json:"forwardedIp" yaml:"forwardedIp"`
	// header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.20.1/docs/resources/wafv2_rule_group#header Wafv2RuleGroup#header}
	Header *Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyHeader `field:"optional" json:"header" yaml:"header"`
	// http_method block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.20.1/docs/resources/wafv2_rule_group#http_method Wafv2RuleGroup#http_method}
	HttpMethod *Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyHttpMethod `field:"optional" json:"httpMethod" yaml:"httpMethod"`
	// ip block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.20.1/docs/resources/wafv2_rule_group#ip Wafv2RuleGroup#ip}
	Ip *Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyIp `field:"optional" json:"ip" yaml:"ip"`
	// label_namespace block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.20.1/docs/resources/wafv2_rule_group#label_namespace Wafv2RuleGroup#label_namespace}
	LabelNamespace *Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyLabelNamespace `field:"optional" json:"labelNamespace" yaml:"labelNamespace"`
	// query_argument block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.20.1/docs/resources/wafv2_rule_group#query_argument Wafv2RuleGroup#query_argument}
	QueryArgument *Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyQueryArgument `field:"optional" json:"queryArgument" yaml:"queryArgument"`
	// query_string block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.20.1/docs/resources/wafv2_rule_group#query_string Wafv2RuleGroup#query_string}
	QueryString *Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyQueryString `field:"optional" json:"queryString" yaml:"queryString"`
	// uri_path block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.20.1/docs/resources/wafv2_rule_group#uri_path Wafv2RuleGroup#uri_path}
	UriPath *Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyUriPath `field:"optional" json:"uriPath" yaml:"uriPath"`
}

