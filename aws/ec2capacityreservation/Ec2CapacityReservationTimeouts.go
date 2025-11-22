// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2capacityreservation


type Ec2CapacityReservationTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/ec2_capacity_reservation#create Ec2CapacityReservation#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/ec2_capacity_reservation#delete Ec2CapacityReservation#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/ec2_capacity_reservation#update Ec2CapacityReservation#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

