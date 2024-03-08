// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ebssnapshotcopy


type EbsSnapshotCopyTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.40.0/docs/resources/ebs_snapshot_copy#create EbsSnapshotCopy#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.40.0/docs/resources/ebs_snapshot_copy#delete EbsSnapshotCopy#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

