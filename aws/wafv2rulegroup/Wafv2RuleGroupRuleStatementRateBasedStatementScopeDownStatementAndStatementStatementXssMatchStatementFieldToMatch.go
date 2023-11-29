// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafv2rulegroup


type Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementXssMatchStatementFieldToMatch struct {
	// all_query_arguments block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/wafv2_rule_group#all_query_arguments Wafv2RuleGroup#all_query_arguments}
	AllQueryArguments *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementXssMatchStatementFieldToMatchAllQueryArguments `field:"optional" json:"allQueryArguments" yaml:"allQueryArguments"`
	// body block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/wafv2_rule_group#body Wafv2RuleGroup#body}
	Body *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementXssMatchStatementFieldToMatchBody `field:"optional" json:"body" yaml:"body"`
	// cookies block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/wafv2_rule_group#cookies Wafv2RuleGroup#cookies}
	Cookies *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementXssMatchStatementFieldToMatchCookies `field:"optional" json:"cookies" yaml:"cookies"`
	// headers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/wafv2_rule_group#headers Wafv2RuleGroup#headers}
	Headers interface{} `field:"optional" json:"headers" yaml:"headers"`
	// ja3_fingerprint block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/wafv2_rule_group#ja3_fingerprint Wafv2RuleGroup#ja3_fingerprint}
	Ja3Fingerprint *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementXssMatchStatementFieldToMatchJa3Fingerprint `field:"optional" json:"ja3Fingerprint" yaml:"ja3Fingerprint"`
	// json_body block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/wafv2_rule_group#json_body Wafv2RuleGroup#json_body}
	JsonBody *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementXssMatchStatementFieldToMatchJsonBody `field:"optional" json:"jsonBody" yaml:"jsonBody"`
	// method block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/wafv2_rule_group#method Wafv2RuleGroup#method}
	Method *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementXssMatchStatementFieldToMatchMethod `field:"optional" json:"method" yaml:"method"`
	// query_string block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/wafv2_rule_group#query_string Wafv2RuleGroup#query_string}
	QueryString *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementXssMatchStatementFieldToMatchQueryString `field:"optional" json:"queryString" yaml:"queryString"`
	// single_header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/wafv2_rule_group#single_header Wafv2RuleGroup#single_header}
	SingleHeader *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementXssMatchStatementFieldToMatchSingleHeader `field:"optional" json:"singleHeader" yaml:"singleHeader"`
	// single_query_argument block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/wafv2_rule_group#single_query_argument Wafv2RuleGroup#single_query_argument}
	SingleQueryArgument *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementXssMatchStatementFieldToMatchSingleQueryArgument `field:"optional" json:"singleQueryArgument" yaml:"singleQueryArgument"`
	// uri_path block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/wafv2_rule_group#uri_path Wafv2RuleGroup#uri_path}
	UriPath *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementXssMatchStatementFieldToMatchUriPath `field:"optional" json:"uriPath" yaml:"uriPath"`
}

