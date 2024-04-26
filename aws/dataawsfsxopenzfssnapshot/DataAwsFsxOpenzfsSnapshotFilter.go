// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsfsxopenzfssnapshot


type DataAwsFsxOpenzfsSnapshotFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/data-sources/fsx_openzfs_snapshot#name DataAwsFsxOpenzfsSnapshot#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/data-sources/fsx_openzfs_snapshot#values DataAwsFsxOpenzfsSnapshot#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

