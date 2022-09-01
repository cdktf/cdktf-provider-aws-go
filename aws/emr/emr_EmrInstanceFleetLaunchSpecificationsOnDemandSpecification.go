package emr


type EmrInstanceFleetLaunchSpecificationsOnDemandSpecification struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_instance_fleet#allocation_strategy EmrInstanceFleet#allocation_strategy}.
	AllocationStrategy *string `field:"required" json:"allocationStrategy" yaml:"allocationStrategy"`
}

