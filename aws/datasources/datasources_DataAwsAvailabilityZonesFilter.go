package datasources


type DataAwsAvailabilityZonesFilter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/availability_zones#name DataAwsAvailabilityZones#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/availability_zones#values DataAwsAvailabilityZones#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

