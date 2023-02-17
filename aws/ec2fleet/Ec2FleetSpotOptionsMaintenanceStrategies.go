package ec2fleet


type Ec2FleetSpotOptionsMaintenanceStrategies struct {
	// capacity_rebalance block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ec2_fleet#capacity_rebalance Ec2Fleet#capacity_rebalance}
	CapacityRebalance *Ec2FleetSpotOptionsMaintenanceStrategiesCapacityRebalance `field:"optional" json:"capacityRebalance" yaml:"capacityRebalance"`
}

