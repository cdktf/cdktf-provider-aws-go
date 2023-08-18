package wafv2webacl


type Wafv2WebAclRuleStatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesAtpRuleSetResponseInspection struct {
	// body_contains block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/wafv2_web_acl#body_contains Wafv2WebAcl#body_contains}
	BodyContains *Wafv2WebAclRuleStatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesAtpRuleSetResponseInspectionBodyContains `field:"optional" json:"bodyContains" yaml:"bodyContains"`
	// header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/wafv2_web_acl#header Wafv2WebAcl#header}
	Header *Wafv2WebAclRuleStatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesAtpRuleSetResponseInspectionHeader `field:"optional" json:"header" yaml:"header"`
	// json block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/wafv2_web_acl#json Wafv2WebAcl#json}
	Json *Wafv2WebAclRuleStatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesAtpRuleSetResponseInspectionJson `field:"optional" json:"json" yaml:"json"`
	// status_code block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/wafv2_web_acl#status_code Wafv2WebAcl#status_code}
	StatusCode *Wafv2WebAclRuleStatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesAtpRuleSetResponseInspectionStatusCode `field:"optional" json:"statusCode" yaml:"statusCode"`
}

