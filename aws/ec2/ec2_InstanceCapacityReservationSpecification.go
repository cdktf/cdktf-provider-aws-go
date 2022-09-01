package ec2


type InstanceCapacityReservationSpecification struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/instance#capacity_reservation_preference Instance#capacity_reservation_preference}.
	CapacityReservationPreference *string `field:"optional" json:"capacityReservationPreference" yaml:"capacityReservationPreference"`
	// capacity_reservation_target block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/instance#capacity_reservation_target Instance#capacity_reservation_target}
	CapacityReservationTarget *InstanceCapacityReservationSpecificationCapacityReservationTarget `field:"optional" json:"capacityReservationTarget" yaml:"capacityReservationTarget"`
}

