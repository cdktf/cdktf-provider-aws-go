package vpc


type DefaultSubnetTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/default_subnet#create DefaultSubnet#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/default_subnet#delete DefaultSubnet#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

