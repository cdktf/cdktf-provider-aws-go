// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package spotfleetrequest


type SpotFleetRequestSpotMaintenanceStrategiesCapacityRebalance struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/spot_fleet_request#replacement_strategy SpotFleetRequest#replacement_strategy}.
	ReplacementStrategy *string `field:"optional" json:"replacementStrategy" yaml:"replacementStrategy"`
}

