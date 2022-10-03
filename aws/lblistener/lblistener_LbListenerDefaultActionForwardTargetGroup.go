package lblistener


type LbListenerDefaultActionForwardTargetGroup struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lb_listener#arn LbListener#arn}.
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lb_listener#weight LbListener#weight}.
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

