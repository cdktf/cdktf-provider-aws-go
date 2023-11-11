// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafv2webacl


type Wafv2WebAclRuleStatementRateBasedStatementCustomKey struct {
	// cookie block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.25.0/docs/resources/wafv2_web_acl#cookie Wafv2WebAcl#cookie}
	Cookie *Wafv2WebAclRuleStatementRateBasedStatementCustomKeyCookie `field:"optional" json:"cookie" yaml:"cookie"`
	// forwarded_ip block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.25.0/docs/resources/wafv2_web_acl#forwarded_ip Wafv2WebAcl#forwarded_ip}
	ForwardedIp *Wafv2WebAclRuleStatementRateBasedStatementCustomKeyForwardedIp `field:"optional" json:"forwardedIp" yaml:"forwardedIp"`
	// header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.25.0/docs/resources/wafv2_web_acl#header Wafv2WebAcl#header}
	Header *Wafv2WebAclRuleStatementRateBasedStatementCustomKeyHeader `field:"optional" json:"header" yaml:"header"`
	// http_method block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.25.0/docs/resources/wafv2_web_acl#http_method Wafv2WebAcl#http_method}
	HttpMethod *Wafv2WebAclRuleStatementRateBasedStatementCustomKeyHttpMethod `field:"optional" json:"httpMethod" yaml:"httpMethod"`
	// ip block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.25.0/docs/resources/wafv2_web_acl#ip Wafv2WebAcl#ip}
	Ip *Wafv2WebAclRuleStatementRateBasedStatementCustomKeyIp `field:"optional" json:"ip" yaml:"ip"`
	// label_namespace block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.25.0/docs/resources/wafv2_web_acl#label_namespace Wafv2WebAcl#label_namespace}
	LabelNamespace *Wafv2WebAclRuleStatementRateBasedStatementCustomKeyLabelNamespace `field:"optional" json:"labelNamespace" yaml:"labelNamespace"`
	// query_argument block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.25.0/docs/resources/wafv2_web_acl#query_argument Wafv2WebAcl#query_argument}
	QueryArgument *Wafv2WebAclRuleStatementRateBasedStatementCustomKeyQueryArgument `field:"optional" json:"queryArgument" yaml:"queryArgument"`
	// query_string block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.25.0/docs/resources/wafv2_web_acl#query_string Wafv2WebAcl#query_string}
	QueryString *Wafv2WebAclRuleStatementRateBasedStatementCustomKeyQueryString `field:"optional" json:"queryString" yaml:"queryString"`
	// uri_path block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.25.0/docs/resources/wafv2_web_acl#uri_path Wafv2WebAcl#uri_path}
	UriPath *Wafv2WebAclRuleStatementRateBasedStatementCustomKeyUriPath `field:"optional" json:"uriPath" yaml:"uriPath"`
}

