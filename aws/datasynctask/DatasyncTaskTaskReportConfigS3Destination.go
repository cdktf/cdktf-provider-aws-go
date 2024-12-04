// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasynctask


type DatasyncTaskTaskReportConfigS3Destination struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.79.0/docs/resources/datasync_task#bucket_access_role_arn DatasyncTask#bucket_access_role_arn}.
	BucketAccessRoleArn *string `field:"required" json:"bucketAccessRoleArn" yaml:"bucketAccessRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.79.0/docs/resources/datasync_task#s3_bucket_arn DatasyncTask#s3_bucket_arn}.
	S3BucketArn *string `field:"required" json:"s3BucketArn" yaml:"s3BucketArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.79.0/docs/resources/datasync_task#subdirectory DatasyncTask#subdirectory}.
	Subdirectory *string `field:"optional" json:"subdirectory" yaml:"subdirectory"`
}

