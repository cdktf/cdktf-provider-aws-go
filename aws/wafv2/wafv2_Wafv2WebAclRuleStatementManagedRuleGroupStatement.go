package wafv2


type Wafv2WebAclRuleStatementManagedRuleGroupStatement struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/wafv2_web_acl#name Wafv2WebAcl#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/wafv2_web_acl#vendor_name Wafv2WebAcl#vendor_name}.
	VendorName *string `field:"required" json:"vendorName" yaml:"vendorName"`
	// excluded_rule block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/wafv2_web_acl#excluded_rule Wafv2WebAcl#excluded_rule}
	ExcludedRule interface{} `field:"optional" json:"excludedRule" yaml:"excludedRule"`
	// scope_down_statement block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/wafv2_web_acl#scope_down_statement Wafv2WebAcl#scope_down_statement}
	ScopeDownStatement *Wafv2WebAclRuleStatementManagedRuleGroupStatementScopeDownStatement `field:"optional" json:"scopeDownStatement" yaml:"scopeDownStatement"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/wafv2_web_acl#version Wafv2WebAcl#version}.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

