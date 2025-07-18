// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ebssnapshot


type EbsSnapshotTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/ebs_snapshot#create EbsSnapshot#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/ebs_snapshot#delete EbsSnapshot#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

