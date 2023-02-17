package alblistener


type AlbListenerDefaultActionForwardStickiness struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/alb_listener#duration AlbListener#duration}.
	Duration *float64 `field:"required" json:"duration" yaml:"duration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/alb_listener#enabled AlbListener#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

