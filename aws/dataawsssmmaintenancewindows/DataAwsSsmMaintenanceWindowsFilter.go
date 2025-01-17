// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsssmmaintenancewindows


type DataAwsSsmMaintenanceWindowsFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/data-sources/ssm_maintenance_windows#name DataAwsSsmMaintenanceWindows#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/data-sources/ssm_maintenance_windows#values DataAwsSsmMaintenanceWindows#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

