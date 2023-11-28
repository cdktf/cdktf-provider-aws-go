// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ebssnapshotimport


type EbsSnapshotImportDiskContainerUserBucket struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.27.0/docs/resources/ebs_snapshot_import#s3_bucket EbsSnapshotImport#s3_bucket}.
	S3Bucket *string `field:"required" json:"s3Bucket" yaml:"s3Bucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.27.0/docs/resources/ebs_snapshot_import#s3_key EbsSnapshotImport#s3_key}.
	S3Key *string `field:"required" json:"s3Key" yaml:"s3Key"`
}

