package spotfleetrequest


type SpotFleetRequestSpotMaintenanceStrategies struct {
	// capacity_rebalance block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.67.0/docs/resources/spot_fleet_request#capacity_rebalance SpotFleetRequest#capacity_rebalance}
	CapacityRebalance *SpotFleetRequestSpotMaintenanceStrategiesCapacityRebalance `field:"optional" json:"capacityRebalance" yaml:"capacityRebalance"`
}

