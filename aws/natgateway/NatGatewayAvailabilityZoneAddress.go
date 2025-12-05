// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package natgateway


type NatGatewayAvailabilityZoneAddress struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/nat_gateway#allocation_ids NatGateway#allocation_ids}.
	AllocationIds *[]*string `field:"optional" json:"allocationIds" yaml:"allocationIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/nat_gateway#availability_zone NatGateway#availability_zone}.
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/nat_gateway#availability_zone_id NatGateway#availability_zone_id}.
	AvailabilityZoneId *string `field:"optional" json:"availabilityZoneId" yaml:"availabilityZoneId"`
}

