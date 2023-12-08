// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package emrinstancefleet


type EmrInstanceFleetLaunchSpecificationsSpotSpecification struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.30.0/docs/resources/emr_instance_fleet#allocation_strategy EmrInstanceFleet#allocation_strategy}.
	AllocationStrategy *string `field:"required" json:"allocationStrategy" yaml:"allocationStrategy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.30.0/docs/resources/emr_instance_fleet#timeout_action EmrInstanceFleet#timeout_action}.
	TimeoutAction *string `field:"required" json:"timeoutAction" yaml:"timeoutAction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.30.0/docs/resources/emr_instance_fleet#timeout_duration_minutes EmrInstanceFleet#timeout_duration_minutes}.
	TimeoutDurationMinutes *float64 `field:"required" json:"timeoutDurationMinutes" yaml:"timeoutDurationMinutes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.30.0/docs/resources/emr_instance_fleet#block_duration_minutes EmrInstanceFleet#block_duration_minutes}.
	BlockDurationMinutes *float64 `field:"optional" json:"blockDurationMinutes" yaml:"blockDurationMinutes"`
}

