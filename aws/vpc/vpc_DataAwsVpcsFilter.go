package vpc


type DataAwsVpcsFilter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/vpcs#name DataAwsVpcs#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/vpcs#values DataAwsVpcs#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

