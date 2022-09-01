package elb


type AlbListenerRuleActionForwardTargetGroup struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/alb_listener_rule#arn AlbListenerRule#arn}.
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/alb_listener_rule#weight AlbListenerRule#weight}.
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

