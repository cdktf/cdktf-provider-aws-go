// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dbsnapshotcopy


type DbSnapshotCopyTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.0/docs/resources/db_snapshot_copy#create DbSnapshotCopy#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
}

