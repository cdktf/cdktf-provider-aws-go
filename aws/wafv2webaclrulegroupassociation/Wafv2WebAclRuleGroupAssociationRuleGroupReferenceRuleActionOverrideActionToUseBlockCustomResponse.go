// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafv2webaclrulegroupassociation


type Wafv2WebAclRuleGroupAssociationRuleGroupReferenceRuleActionOverrideActionToUseBlockCustomResponse struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/wafv2_web_acl_rule_group_association#response_code Wafv2WebAclRuleGroupAssociation#response_code}.
	ResponseCode *float64 `field:"required" json:"responseCode" yaml:"responseCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/wafv2_web_acl_rule_group_association#custom_response_body_key Wafv2WebAclRuleGroupAssociation#custom_response_body_key}.
	CustomResponseBodyKey *string `field:"optional" json:"customResponseBodyKey" yaml:"customResponseBodyKey"`
	// response_header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/wafv2_web_acl_rule_group_association#response_header Wafv2WebAclRuleGroupAssociation#response_header}
	ResponseHeader interface{} `field:"optional" json:"responseHeader" yaml:"responseHeader"`
}

