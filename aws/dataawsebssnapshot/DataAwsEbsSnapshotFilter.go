// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsebssnapshot


type DataAwsEbsSnapshotFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/data-sources/ebs_snapshot#name DataAwsEbsSnapshot#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/data-sources/ebs_snapshot#values DataAwsEbsSnapshot#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

