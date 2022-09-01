package ec2


type AmiFromInstanceTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ami_from_instance#create AmiFromInstance#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ami_from_instance#delete AmiFromInstance#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ami_from_instance#update AmiFromInstance#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

