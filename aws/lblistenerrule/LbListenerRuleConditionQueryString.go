package lblistenerrule


type LbListenerRuleConditionQueryString struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/lb_listener_rule#value LbListenerRule#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/lb_listener_rule#key LbListenerRule#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
}

