// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package efsaccesspoint


type EfsAccessPointRootDirectoryCreationInfo struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.36.0/docs/resources/efs_access_point#owner_gid EfsAccessPoint#owner_gid}.
	OwnerGid *float64 `field:"required" json:"ownerGid" yaml:"ownerGid"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.36.0/docs/resources/efs_access_point#owner_uid EfsAccessPoint#owner_uid}.
	OwnerUid *float64 `field:"required" json:"ownerUid" yaml:"ownerUid"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.36.0/docs/resources/efs_access_point#permissions EfsAccessPoint#permissions}.
	Permissions *string `field:"required" json:"permissions" yaml:"permissions"`
}

