package datasources


type DataAwsAvailabilityZoneFilter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/availability_zone#name DataAwsAvailabilityZone#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/availability_zone#values DataAwsAvailabilityZone#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

