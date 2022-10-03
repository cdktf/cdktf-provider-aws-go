package lblistener


type LbListenerDefaultActionForward struct {
	// target_group block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lb_listener#target_group LbListener#target_group}
	TargetGroup interface{} `field:"required" json:"targetGroup" yaml:"targetGroup"`
	// stickiness block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lb_listener#stickiness LbListener#stickiness}
	Stickiness *LbListenerDefaultActionForwardStickiness `field:"optional" json:"stickiness" yaml:"stickiness"`
}

