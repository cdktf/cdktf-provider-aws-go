package instance


type InstanceCapacityReservationSpecificationCapacityReservationTarget struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/instance#capacity_reservation_id Instance#capacity_reservation_id}.
	CapacityReservationId *string `field:"optional" json:"capacityReservationId" yaml:"capacityReservationId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/instance#capacity_reservation_resource_group_arn Instance#capacity_reservation_resource_group_arn}.
	CapacityReservationResourceGroupArn *string `field:"optional" json:"capacityReservationResourceGroupArn" yaml:"capacityReservationResourceGroupArn"`
}

