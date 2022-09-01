package elb


type LbListenerRuleActionForward struct {
	// target_group block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lb_listener_rule#target_group LbListenerRule#target_group}
	TargetGroup interface{} `field:"required" json:"targetGroup" yaml:"targetGroup"`
	// stickiness block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lb_listener_rule#stickiness LbListenerRule#stickiness}
	Stickiness *LbListenerRuleActionForwardStickiness `field:"optional" json:"stickiness" yaml:"stickiness"`
}

