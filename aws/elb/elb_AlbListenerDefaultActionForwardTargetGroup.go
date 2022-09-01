package elb


type AlbListenerDefaultActionForwardTargetGroup struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/alb_listener#arn AlbListener#arn}.
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/alb_listener#weight AlbListener#weight}.
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

