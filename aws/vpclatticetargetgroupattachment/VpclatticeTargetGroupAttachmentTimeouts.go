// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpclatticetargetgroupattachment


type VpclatticeTargetGroupAttachmentTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/vpclattice_target_group_attachment#create VpclatticeTargetGroupAttachment#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/vpclattice_target_group_attachment#delete VpclatticeTargetGroupAttachment#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

