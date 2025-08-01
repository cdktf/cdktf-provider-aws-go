// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package emrcluster


type EmrClusterMasterInstanceFleetLaunchSpecificationsSpotSpecification struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/emr_cluster#allocation_strategy EmrCluster#allocation_strategy}.
	AllocationStrategy *string `field:"required" json:"allocationStrategy" yaml:"allocationStrategy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/emr_cluster#timeout_action EmrCluster#timeout_action}.
	TimeoutAction *string `field:"required" json:"timeoutAction" yaml:"timeoutAction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/emr_cluster#timeout_duration_minutes EmrCluster#timeout_duration_minutes}.
	TimeoutDurationMinutes *float64 `field:"required" json:"timeoutDurationMinutes" yaml:"timeoutDurationMinutes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/emr_cluster#block_duration_minutes EmrCluster#block_duration_minutes}.
	BlockDurationMinutes *float64 `field:"optional" json:"blockDurationMinutes" yaml:"blockDurationMinutes"`
}

