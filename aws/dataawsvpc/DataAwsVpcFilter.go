package dataawsvpc


type DataAwsVpcFilter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/vpc#name DataAwsVpc#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/vpc#values DataAwsVpc#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

