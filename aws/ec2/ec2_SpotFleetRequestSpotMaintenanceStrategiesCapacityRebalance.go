package ec2


type SpotFleetRequestSpotMaintenanceStrategiesCapacityRebalance struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/spot_fleet_request#replacement_strategy SpotFleetRequest#replacement_strategy}.
	ReplacementStrategy *string `field:"optional" json:"replacementStrategy" yaml:"replacementStrategy"`
}

