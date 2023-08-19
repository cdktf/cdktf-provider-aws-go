package vpclatticelistenerrule


type VpclatticeListenerRuleMatchHttpMatchHeaderMatchesMatch struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/vpclattice_listener_rule#contains VpclatticeListenerRule#contains}.
	Contains *string `field:"optional" json:"contains" yaml:"contains"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/vpclattice_listener_rule#exact VpclatticeListenerRule#exact}.
	Exact *string `field:"optional" json:"exact" yaml:"exact"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/vpclattice_listener_rule#prefix VpclatticeListenerRule#prefix}.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

