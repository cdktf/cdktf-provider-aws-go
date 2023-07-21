package wafv2webacl


type Wafv2WebAclRuleStatementRateBasedStatement struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/wafv2_web_acl#limit Wafv2WebAcl#limit}.
	Limit *float64 `field:"required" json:"limit" yaml:"limit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/wafv2_web_acl#aggregate_key_type Wafv2WebAcl#aggregate_key_type}.
	AggregateKeyType *string `field:"optional" json:"aggregateKeyType" yaml:"aggregateKeyType"`
	// forwarded_ip_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/wafv2_web_acl#forwarded_ip_config Wafv2WebAcl#forwarded_ip_config}
	ForwardedIpConfig *Wafv2WebAclRuleStatementRateBasedStatementForwardedIpConfig `field:"optional" json:"forwardedIpConfig" yaml:"forwardedIpConfig"`
	// scope_down_statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/wafv2_web_acl#scope_down_statement Wafv2WebAcl#scope_down_statement}
	ScopeDownStatement *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatement `field:"optional" json:"scopeDownStatement" yaml:"scopeDownStatement"`
}

