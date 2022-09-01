package vpc


type SubnetTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/subnet#create Subnet#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/subnet#delete Subnet#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

