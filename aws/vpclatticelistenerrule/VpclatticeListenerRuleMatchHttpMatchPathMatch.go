package vpclatticelistenerrule


type VpclatticeListenerRuleMatchHttpMatchPathMatch struct {
	// match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.8.0/docs/resources/vpclattice_listener_rule#match VpclatticeListenerRule#match}
	Match *VpclatticeListenerRuleMatchHttpMatchPathMatchMatch `field:"required" json:"match" yaml:"match"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.8.0/docs/resources/vpclattice_listener_rule#case_sensitive VpclatticeListenerRule#case_sensitive}.
	CaseSensitive interface{} `field:"optional" json:"caseSensitive" yaml:"caseSensitive"`
}

