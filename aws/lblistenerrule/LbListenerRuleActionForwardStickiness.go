package lblistenerrule


type LbListenerRuleActionForwardStickiness struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/lb_listener_rule#duration LbListenerRule#duration}.
	Duration *float64 `field:"required" json:"duration" yaml:"duration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/lb_listener_rule#enabled LbListenerRule#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

