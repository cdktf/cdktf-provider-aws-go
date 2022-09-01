package ram


type DataAwsRamResourceShareFilter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/ram_resource_share#name DataAwsRamResourceShare#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/ram_resource_share#values DataAwsRamResourceShare#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

