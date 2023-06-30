package ec2fleet


type Ec2FleetSpotOptionsMaintenanceStrategies struct {
	// capacity_rebalance block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.1/docs/resources/ec2_fleet#capacity_rebalance Ec2Fleet#capacity_rebalance}
	CapacityRebalance *Ec2FleetSpotOptionsMaintenanceStrategiesCapacityRebalance `field:"optional" json:"capacityRebalance" yaml:"capacityRebalance"`
}

