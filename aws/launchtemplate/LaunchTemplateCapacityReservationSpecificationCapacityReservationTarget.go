// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package launchtemplate


type LaunchTemplateCapacityReservationSpecificationCapacityReservationTarget struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/launch_template#capacity_reservation_id LaunchTemplate#capacity_reservation_id}.
	CapacityReservationId *string `field:"optional" json:"capacityReservationId" yaml:"capacityReservationId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/launch_template#capacity_reservation_resource_group_arn LaunchTemplate#capacity_reservation_resource_group_arn}.
	CapacityReservationResourceGroupArn *string `field:"optional" json:"capacityReservationResourceGroupArn" yaml:"capacityReservationResourceGroupArn"`
}

