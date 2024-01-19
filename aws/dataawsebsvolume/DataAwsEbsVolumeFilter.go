// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsebsvolume


type DataAwsEbsVolumeFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/data-sources/ebs_volume#name DataAwsEbsVolume#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/data-sources/ebs_volume#values DataAwsEbsVolume#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

