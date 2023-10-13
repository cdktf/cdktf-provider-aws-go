// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ssmmaintenancewindowtarget


type SsmMaintenanceWindowTargetTargets struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.21.0/docs/resources/ssm_maintenance_window_target#key SsmMaintenanceWindowTarget#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.21.0/docs/resources/ssm_maintenance_window_target#values SsmMaintenanceWindowTarget#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

