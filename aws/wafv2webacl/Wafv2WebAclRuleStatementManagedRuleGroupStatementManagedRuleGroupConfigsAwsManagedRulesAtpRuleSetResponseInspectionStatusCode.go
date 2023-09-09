// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafv2webacl


type Wafv2WebAclRuleStatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesAtpRuleSetResponseInspectionStatusCode struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.1/docs/resources/wafv2_web_acl#failure_codes Wafv2WebAcl#failure_codes}.
	FailureCodes *[]*float64 `field:"required" json:"failureCodes" yaml:"failureCodes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.1/docs/resources/wafv2_web_acl#success_codes Wafv2WebAcl#success_codes}.
	SuccessCodes *[]*float64 `field:"required" json:"successCodes" yaml:"successCodes"`
}

