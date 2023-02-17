package dataawsdbinstances


type DataAwsDbInstancesFilter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/db_instances#name DataAwsDbInstances#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/db_instances#values DataAwsDbInstances#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

