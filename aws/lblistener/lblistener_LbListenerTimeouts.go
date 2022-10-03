package lblistener


type LbListenerTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lb_listener#read LbListener#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

