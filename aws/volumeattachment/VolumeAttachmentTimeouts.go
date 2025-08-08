// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package volumeattachment


type VolumeAttachmentTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/volume_attachment#create VolumeAttachment#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/volume_attachment#delete VolumeAttachment#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

