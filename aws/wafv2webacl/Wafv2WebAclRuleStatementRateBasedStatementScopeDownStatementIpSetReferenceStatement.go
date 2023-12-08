// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafv2webacl


type Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementIpSetReferenceStatement struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.30.0/docs/resources/wafv2_web_acl#arn Wafv2WebAcl#arn}.
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// ip_set_forwarded_ip_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.30.0/docs/resources/wafv2_web_acl#ip_set_forwarded_ip_config Wafv2WebAcl#ip_set_forwarded_ip_config}
	IpSetForwardedIpConfig *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementIpSetReferenceStatementIpSetForwardedIpConfig `field:"optional" json:"ipSetForwardedIpConfig" yaml:"ipSetForwardedIpConfig"`
}

