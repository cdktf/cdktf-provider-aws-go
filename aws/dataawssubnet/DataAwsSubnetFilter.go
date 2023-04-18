package dataawssubnet


type DataAwsSubnetFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.63.0/docs/data-sources/subnet#name DataAwsSubnet#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.63.0/docs/data-sources/subnet#values DataAwsSubnet#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

