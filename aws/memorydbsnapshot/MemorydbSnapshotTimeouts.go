// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package memorydbsnapshot


type MemorydbSnapshotTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.70.0/docs/resources/memorydb_snapshot#create MemorydbSnapshot#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.70.0/docs/resources/memorydb_snapshot#delete MemorydbSnapshot#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

