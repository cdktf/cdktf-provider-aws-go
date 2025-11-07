// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package fsxs3accesspointattachment


type FsxS3AccessPointAttachmentS3AccessPoint struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.20.0/docs/resources/fsx_s3_access_point_attachment#policy FsxS3AccessPointAttachment#policy}.
	Policy *string `field:"optional" json:"policy" yaml:"policy"`
	// vpc_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.20.0/docs/resources/fsx_s3_access_point_attachment#vpc_configuration FsxS3AccessPointAttachment#vpc_configuration}
	VpcConfiguration interface{} `field:"optional" json:"vpcConfiguration" yaml:"vpcConfiguration"`
}

