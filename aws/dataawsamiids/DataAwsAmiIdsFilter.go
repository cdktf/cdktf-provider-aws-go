package dataawsamiids


type DataAwsAmiIdsFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/data-sources/ami_ids#name DataAwsAmiIds#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/data-sources/ami_ids#values DataAwsAmiIds#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

