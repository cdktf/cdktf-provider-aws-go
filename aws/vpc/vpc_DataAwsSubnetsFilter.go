package vpc


type DataAwsSubnetsFilter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/subnets#name DataAwsSubnets#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/subnets#values DataAwsSubnets#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

