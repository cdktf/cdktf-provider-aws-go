package ec2


type DataAwsEipFilter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/eip#name DataAwsEip#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/eip#values DataAwsEip#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

