// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package spotinstancerequest


type SpotInstanceRequestCapacityReservationSpecificationCapacityReservationTarget struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/spot_instance_request#capacity_reservation_id SpotInstanceRequest#capacity_reservation_id}.
	CapacityReservationId *string `field:"optional" json:"capacityReservationId" yaml:"capacityReservationId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/spot_instance_request#capacity_reservation_resource_group_arn SpotInstanceRequest#capacity_reservation_resource_group_arn}.
	CapacityReservationResourceGroupArn *string `field:"optional" json:"capacityReservationResourceGroupArn" yaml:"capacityReservationResourceGroupArn"`
}

