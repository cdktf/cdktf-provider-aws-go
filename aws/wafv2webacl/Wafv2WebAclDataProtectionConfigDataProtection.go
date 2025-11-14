// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafv2webacl


type Wafv2WebAclDataProtectionConfigDataProtection struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.21.0/docs/resources/wafv2_web_acl#action Wafv2WebAcl#action}.
	Action *string `field:"required" json:"action" yaml:"action"`
	// field block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.21.0/docs/resources/wafv2_web_acl#field Wafv2WebAcl#field}
	Field *Wafv2WebAclDataProtectionConfigDataProtectionField `field:"required" json:"field" yaml:"field"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.21.0/docs/resources/wafv2_web_acl#exclude_rate_based_details Wafv2WebAcl#exclude_rate_based_details}.
	ExcludeRateBasedDetails interface{} `field:"optional" json:"excludeRateBasedDetails" yaml:"excludeRateBasedDetails"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.21.0/docs/resources/wafv2_web_acl#exclude_rule_match_details Wafv2WebAcl#exclude_rule_match_details}.
	ExcludeRuleMatchDetails interface{} `field:"optional" json:"excludeRuleMatchDetails" yaml:"excludeRuleMatchDetails"`
}

