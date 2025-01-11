// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package efsaccesspoint


type EfsAccessPointPosixUser struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.1/docs/resources/efs_access_point#gid EfsAccessPoint#gid}.
	Gid *float64 `field:"required" json:"gid" yaml:"gid"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.1/docs/resources/efs_access_point#uid EfsAccessPoint#uid}.
	Uid *float64 `field:"required" json:"uid" yaml:"uid"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.1/docs/resources/efs_access_point#secondary_gids EfsAccessPoint#secondary_gids}.
	SecondaryGids *[]*float64 `field:"optional" json:"secondaryGids" yaml:"secondaryGids"`
}

