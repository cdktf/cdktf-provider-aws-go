// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package fsxs3accesspointattachment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FsxS3AccessPointAttachmentConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.20.0/docs/resources/fsx_s3_access_point_attachment#name FsxS3AccessPointAttachment#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.20.0/docs/resources/fsx_s3_access_point_attachment#type FsxS3AccessPointAttachment#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// openzfs_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.20.0/docs/resources/fsx_s3_access_point_attachment#openzfs_configuration FsxS3AccessPointAttachment#openzfs_configuration}
	OpenzfsConfiguration interface{} `field:"optional" json:"openzfsConfiguration" yaml:"openzfsConfiguration"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.20.0/docs/resources/fsx_s3_access_point_attachment#region FsxS3AccessPointAttachment#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// s3_access_point block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.20.0/docs/resources/fsx_s3_access_point_attachment#s3_access_point FsxS3AccessPointAttachment#s3_access_point}
	S3AccessPoint interface{} `field:"optional" json:"s3AccessPoint" yaml:"s3AccessPoint"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.20.0/docs/resources/fsx_s3_access_point_attachment#timeouts FsxS3AccessPointAttachment#timeouts}
	Timeouts *FsxS3AccessPointAttachmentTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

