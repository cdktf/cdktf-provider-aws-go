package alblistener


type AlbListenerTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/alb_listener#read AlbListener#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

