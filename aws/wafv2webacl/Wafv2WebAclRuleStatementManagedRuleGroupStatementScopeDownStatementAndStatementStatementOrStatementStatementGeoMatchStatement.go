// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafv2webacl


type Wafv2WebAclRuleStatementManagedRuleGroupStatementScopeDownStatementAndStatementStatementOrStatementStatementGeoMatchStatement struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/wafv2_web_acl#country_codes Wafv2WebAcl#country_codes}.
	CountryCodes *[]*string `field:"required" json:"countryCodes" yaml:"countryCodes"`
	// forwarded_ip_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/wafv2_web_acl#forwarded_ip_config Wafv2WebAcl#forwarded_ip_config}
	ForwardedIpConfig *Wafv2WebAclRuleStatementManagedRuleGroupStatementScopeDownStatementAndStatementStatementOrStatementStatementGeoMatchStatementForwardedIpConfig `field:"optional" json:"forwardedIpConfig" yaml:"forwardedIpConfig"`
}

