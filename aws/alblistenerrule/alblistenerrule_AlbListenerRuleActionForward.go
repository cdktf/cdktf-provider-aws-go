package alblistenerrule


type AlbListenerRuleActionForward struct {
	// target_group block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/alb_listener_rule#target_group AlbListenerRule#target_group}
	TargetGroup interface{} `field:"required" json:"targetGroup" yaml:"targetGroup"`
	// stickiness block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/alb_listener_rule#stickiness AlbListenerRule#stickiness}
	Stickiness *AlbListenerRuleActionForwardStickiness `field:"optional" json:"stickiness" yaml:"stickiness"`
}

