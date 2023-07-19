package lblistener


type LbListenerDefaultActionForwardTargetGroup struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.8.0/docs/resources/lb_listener#arn LbListener#arn}.
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.8.0/docs/resources/lb_listener#weight LbListener#weight}.
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

