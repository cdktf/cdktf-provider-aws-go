// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ssmmaintenancewindowtask


type SsmMaintenanceWindowTaskTargets struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.1/docs/resources/ssm_maintenance_window_task#key SsmMaintenanceWindowTask#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.1/docs/resources/ssm_maintenance_window_task#values SsmMaintenanceWindowTask#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

