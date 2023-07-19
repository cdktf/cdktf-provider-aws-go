package wafv2webacl


type Wafv2WebAclRuleStatementManagedRuleGroupStatementManagedRuleGroupConfigs struct {
	// aws_managed_rules_atp_rule_set block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.8.0/docs/resources/wafv2_web_acl#aws_managed_rules_atp_rule_set Wafv2WebAcl#aws_managed_rules_atp_rule_set}
	AwsManagedRulesAtpRuleSet *Wafv2WebAclRuleStatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesAtpRuleSet `field:"optional" json:"awsManagedRulesAtpRuleSet" yaml:"awsManagedRulesAtpRuleSet"`
	// aws_managed_rules_bot_control_rule_set block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.8.0/docs/resources/wafv2_web_acl#aws_managed_rules_bot_control_rule_set Wafv2WebAcl#aws_managed_rules_bot_control_rule_set}
	AwsManagedRulesBotControlRuleSet *Wafv2WebAclRuleStatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesBotControlRuleSet `field:"optional" json:"awsManagedRulesBotControlRuleSet" yaml:"awsManagedRulesBotControlRuleSet"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.8.0/docs/resources/wafv2_web_acl#login_path Wafv2WebAcl#login_path}.
	LoginPath *string `field:"optional" json:"loginPath" yaml:"loginPath"`
	// password_field block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.8.0/docs/resources/wafv2_web_acl#password_field Wafv2WebAcl#password_field}
	PasswordField *Wafv2WebAclRuleStatementManagedRuleGroupStatementManagedRuleGroupConfigsPasswordField `field:"optional" json:"passwordField" yaml:"passwordField"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.8.0/docs/resources/wafv2_web_acl#payload_type Wafv2WebAcl#payload_type}.
	PayloadType *string `field:"optional" json:"payloadType" yaml:"payloadType"`
	// username_field block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.8.0/docs/resources/wafv2_web_acl#username_field Wafv2WebAcl#username_field}
	UsernameField *Wafv2WebAclRuleStatementManagedRuleGroupStatementManagedRuleGroupConfigsUsernameField `field:"optional" json:"usernameField" yaml:"usernameField"`
}

