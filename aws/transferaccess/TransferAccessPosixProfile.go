// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package transferaccess


type TransferAccessPosixProfile struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/resources/transfer_access#gid TransferAccess#gid}.
	Gid *float64 `field:"required" json:"gid" yaml:"gid"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/resources/transfer_access#uid TransferAccess#uid}.
	Uid *float64 `field:"required" json:"uid" yaml:"uid"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/resources/transfer_access#secondary_gids TransferAccess#secondary_gids}.
	SecondaryGids *[]*float64 `field:"optional" json:"secondaryGids" yaml:"secondaryGids"`
}

