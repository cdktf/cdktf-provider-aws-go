package emr


type EmrserverlessApplicationInitialCapacity struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emrserverless_application#initial_capacity_type EmrserverlessApplication#initial_capacity_type}.
	InitialCapacityType *string `field:"required" json:"initialCapacityType" yaml:"initialCapacityType"`
	// initial_capacity_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emrserverless_application#initial_capacity_config EmrserverlessApplication#initial_capacity_config}
	InitialCapacityConfig *EmrserverlessApplicationInitialCapacityInitialCapacityConfig `field:"optional" json:"initialCapacityConfig" yaml:"initialCapacityConfig"`
}

