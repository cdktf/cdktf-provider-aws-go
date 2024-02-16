// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package fsxopenzfsvolume


type FsxOpenzfsVolumeNfsExportsClientConfigurations struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.37.0/docs/resources/fsx_openzfs_volume#clients FsxOpenzfsVolume#clients}.
	Clients *string `field:"required" json:"clients" yaml:"clients"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.37.0/docs/resources/fsx_openzfs_volume#options FsxOpenzfsVolume#options}.
	Options *[]*string `field:"required" json:"options" yaml:"options"`
}

