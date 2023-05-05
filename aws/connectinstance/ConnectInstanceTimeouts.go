package connectinstance


type ConnectInstanceTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/connect_instance#create ConnectInstance#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/connect_instance#delete ConnectInstance#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

