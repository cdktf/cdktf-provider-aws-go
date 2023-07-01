package dataawseips


type DataAwsEipsFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/data-sources/eips#name DataAwsEips#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/data-sources/eips#values DataAwsEips#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

