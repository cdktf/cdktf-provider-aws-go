// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package fsxs3accesspointattachment


type FsxS3AccessPointAttachmentOpenzfsConfigurationFileSystemIdentity struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/fsx_s3_access_point_attachment#type FsxS3AccessPointAttachment#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// posix_user block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/fsx_s3_access_point_attachment#posix_user FsxS3AccessPointAttachment#posix_user}
	PosixUser interface{} `field:"optional" json:"posixUser" yaml:"posixUser"`
}

