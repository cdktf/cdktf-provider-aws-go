package datasources


type DataAwsRegionsFilter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/regions#name DataAwsRegions#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/regions#values DataAwsRegions#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

