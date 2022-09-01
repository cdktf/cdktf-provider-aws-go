package ec2


type LaunchTemplateCapacityReservationSpecification struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/launch_template#capacity_reservation_preference LaunchTemplate#capacity_reservation_preference}.
	CapacityReservationPreference *string `field:"optional" json:"capacityReservationPreference" yaml:"capacityReservationPreference"`
	// capacity_reservation_target block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/launch_template#capacity_reservation_target LaunchTemplate#capacity_reservation_target}
	CapacityReservationTarget *LaunchTemplateCapacityReservationSpecificationCapacityReservationTarget `field:"optional" json:"capacityReservationTarget" yaml:"capacityReservationTarget"`
}

