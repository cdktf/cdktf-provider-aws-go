// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsebsvolumes


type DataAwsEbsVolumesFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.86.0/docs/data-sources/ebs_volumes#name DataAwsEbsVolumes#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.86.0/docs/data-sources/ebs_volumes#values DataAwsEbsVolumes#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

