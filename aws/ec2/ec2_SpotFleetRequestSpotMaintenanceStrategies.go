package ec2


type SpotFleetRequestSpotMaintenanceStrategies struct {
	// capacity_rebalance block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/spot_fleet_request#capacity_rebalance SpotFleetRequest#capacity_rebalance}
	CapacityRebalance *SpotFleetRequestSpotMaintenanceStrategiesCapacityRebalance `field:"optional" json:"capacityRebalance" yaml:"capacityRebalance"`
}

