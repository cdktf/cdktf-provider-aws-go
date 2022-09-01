package ec2


type Ec2FleetSpotOptionsMaintenanceStrategiesCapacityRebalance struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ec2_fleet#replacement_strategy Ec2Fleet#replacement_strategy}.
	ReplacementStrategy *string `field:"optional" json:"replacementStrategy" yaml:"replacementStrategy"`
}

