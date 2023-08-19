package wafv2webacl


type Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementGeoMatchStatement struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/wafv2_web_acl#country_codes Wafv2WebAcl#country_codes}.
	CountryCodes *[]*string `field:"required" json:"countryCodes" yaml:"countryCodes"`
	// forwarded_ip_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/wafv2_web_acl#forwarded_ip_config Wafv2WebAcl#forwarded_ip_config}
	ForwardedIpConfig *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementGeoMatchStatementForwardedIpConfig `field:"optional" json:"forwardedIpConfig" yaml:"forwardedIpConfig"`
}

