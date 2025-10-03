// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2fleet


type Ec2FleetOnDemandOptionsCapacityReservationOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/ec2_fleet#usage_strategy Ec2Fleet#usage_strategy}.
	UsageStrategy *string `field:"optional" json:"usageStrategy" yaml:"usageStrategy"`
}

