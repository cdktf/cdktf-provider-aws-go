package appstreamfleet


type AppstreamFleetComputeCapacity struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.8.0/docs/resources/appstream_fleet#desired_instances AppstreamFleet#desired_instances}.
	DesiredInstances *float64 `field:"required" json:"desiredInstances" yaml:"desiredInstances"`
}

