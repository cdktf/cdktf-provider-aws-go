package dataawsvpc


type DataAwsVpcFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/data-sources/vpc#name DataAwsVpc#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/data-sources/vpc#values DataAwsVpc#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

