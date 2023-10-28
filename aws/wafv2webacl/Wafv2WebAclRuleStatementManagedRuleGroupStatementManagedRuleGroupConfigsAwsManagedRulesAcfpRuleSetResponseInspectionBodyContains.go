// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafv2webacl


type Wafv2WebAclRuleStatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesAcfpRuleSetResponseInspectionBodyContains struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.1/docs/resources/wafv2_web_acl#failure_strings Wafv2WebAcl#failure_strings}.
	FailureStrings *[]*string `field:"required" json:"failureStrings" yaml:"failureStrings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.1/docs/resources/wafv2_web_acl#success_strings Wafv2WebAcl#success_strings}.
	SuccessStrings *[]*string `field:"required" json:"successStrings" yaml:"successStrings"`
}

