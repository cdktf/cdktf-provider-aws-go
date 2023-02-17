package dataawssubnetids


type DataAwsSubnetIdsFilter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/subnet_ids#name DataAwsSubnetIds#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/subnet_ids#values DataAwsSubnetIds#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

