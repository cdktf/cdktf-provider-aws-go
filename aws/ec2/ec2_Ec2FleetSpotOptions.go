package ec2


type Ec2FleetSpotOptions struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ec2_fleet#allocation_strategy Ec2Fleet#allocation_strategy}.
	AllocationStrategy *string `field:"optional" json:"allocationStrategy" yaml:"allocationStrategy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ec2_fleet#instance_interruption_behavior Ec2Fleet#instance_interruption_behavior}.
	InstanceInterruptionBehavior *string `field:"optional" json:"instanceInterruptionBehavior" yaml:"instanceInterruptionBehavior"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ec2_fleet#instance_pools_to_use_count Ec2Fleet#instance_pools_to_use_count}.
	InstancePoolsToUseCount *float64 `field:"optional" json:"instancePoolsToUseCount" yaml:"instancePoolsToUseCount"`
	// maintenance_strategies block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ec2_fleet#maintenance_strategies Ec2Fleet#maintenance_strategies}
	MaintenanceStrategies *Ec2FleetSpotOptionsMaintenanceStrategies `field:"optional" json:"maintenanceStrategies" yaml:"maintenanceStrategies"`
}

