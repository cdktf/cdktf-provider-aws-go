// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafv2webacl


type Wafv2WebAclRuleStatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesAcfpRuleSetRequestInspection struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.1/docs/resources/wafv2_web_acl#payload_type Wafv2WebAcl#payload_type}.
	PayloadType *string `field:"required" json:"payloadType" yaml:"payloadType"`
	// email_field block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.1/docs/resources/wafv2_web_acl#email_field Wafv2WebAcl#email_field}
	EmailField *Wafv2WebAclRuleStatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesAcfpRuleSetRequestInspectionEmailField `field:"optional" json:"emailField" yaml:"emailField"`
	// password_field block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.1/docs/resources/wafv2_web_acl#password_field Wafv2WebAcl#password_field}
	PasswordField *Wafv2WebAclRuleStatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesAcfpRuleSetRequestInspectionPasswordField `field:"optional" json:"passwordField" yaml:"passwordField"`
	// username_field block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.1/docs/resources/wafv2_web_acl#username_field Wafv2WebAcl#username_field}
	UsernameField *Wafv2WebAclRuleStatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesAcfpRuleSetRequestInspectionUsernameField `field:"optional" json:"usernameField" yaml:"usernameField"`
}

