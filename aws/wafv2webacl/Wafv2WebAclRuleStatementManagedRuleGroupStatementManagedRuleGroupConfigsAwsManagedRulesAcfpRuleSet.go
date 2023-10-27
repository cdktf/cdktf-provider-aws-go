// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafv2webacl


type Wafv2WebAclRuleStatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesAcfpRuleSet struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/wafv2_web_acl#creation_path Wafv2WebAcl#creation_path}.
	CreationPath *string `field:"required" json:"creationPath" yaml:"creationPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/wafv2_web_acl#registration_page_path Wafv2WebAcl#registration_page_path}.
	RegistrationPagePath *string `field:"required" json:"registrationPagePath" yaml:"registrationPagePath"`
	// request_inspection block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/wafv2_web_acl#request_inspection Wafv2WebAcl#request_inspection}
	RequestInspection *Wafv2WebAclRuleStatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesAcfpRuleSetRequestInspection `field:"required" json:"requestInspection" yaml:"requestInspection"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/wafv2_web_acl#enable_regex_in_path Wafv2WebAcl#enable_regex_in_path}.
	EnableRegexInPath interface{} `field:"optional" json:"enableRegexInPath" yaml:"enableRegexInPath"`
	// response_inspection block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/wafv2_web_acl#response_inspection Wafv2WebAcl#response_inspection}
	ResponseInspection *Wafv2WebAclRuleStatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesAcfpRuleSetResponseInspection `field:"optional" json:"responseInspection" yaml:"responseInspection"`
}

