package connectphonenumber


type ConnectPhoneNumberTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/resources/connect_phone_number#create ConnectPhoneNumber#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/resources/connect_phone_number#delete ConnectPhoneNumber#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/resources/connect_phone_number#update ConnectPhoneNumber#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

