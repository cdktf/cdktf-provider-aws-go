package appstream


type AppstreamFleetComputeCapacity struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_fleet#desired_instances AppstreamFleet#desired_instances}.
	DesiredInstances *float64 `field:"required" json:"desiredInstances" yaml:"desiredInstances"`
}

