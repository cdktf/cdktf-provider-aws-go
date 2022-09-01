package elb


type LbListenerDefaultActionForwardStickiness struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lb_listener#duration LbListener#duration}.
	Duration *float64 `field:"required" json:"duration" yaml:"duration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lb_listener#enabled LbListener#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

