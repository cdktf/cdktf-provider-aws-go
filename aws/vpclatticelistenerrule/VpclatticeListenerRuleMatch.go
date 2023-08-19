package vpclatticelistenerrule


type VpclatticeListenerRuleMatch struct {
	// http_match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/vpclattice_listener_rule#http_match VpclatticeListenerRule#http_match}
	HttpMatch *VpclatticeListenerRuleMatchHttpMatch `field:"optional" json:"httpMatch" yaml:"httpMatch"`
}

