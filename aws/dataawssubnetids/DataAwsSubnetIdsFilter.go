package dataawssubnetids


type DataAwsSubnetIdsFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/data-sources/subnet_ids#name DataAwsSubnetIds#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/data-sources/subnet_ids#values DataAwsSubnetIds#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

