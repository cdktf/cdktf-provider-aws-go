package wafv2webacl


type Wafv2WebAclRuleStatementManagedRuleGroupStatementScopeDownStatementIpSetReferenceStatementIpSetForwardedIpConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/wafv2_web_acl#fallback_behavior Wafv2WebAcl#fallback_behavior}.
	FallbackBehavior *string `field:"required" json:"fallbackBehavior" yaml:"fallbackBehavior"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/wafv2_web_acl#header_name Wafv2WebAcl#header_name}.
	HeaderName *string `field:"required" json:"headerName" yaml:"headerName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/wafv2_web_acl#position Wafv2WebAcl#position}.
	Position *string `field:"required" json:"position" yaml:"position"`
}

