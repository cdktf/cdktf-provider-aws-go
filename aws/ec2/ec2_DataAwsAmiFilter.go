package ec2


type DataAwsAmiFilter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/ami#name DataAwsAmi#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/ami#values DataAwsAmi#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

