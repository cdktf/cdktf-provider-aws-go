// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ebssnapshotimport


type EbsSnapshotImportDiskContainer struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/ebs_snapshot_import#format EbsSnapshotImport#format}.
	Format *string `field:"required" json:"format" yaml:"format"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/ebs_snapshot_import#description EbsSnapshotImport#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/ebs_snapshot_import#url EbsSnapshotImport#url}.
	Url *string `field:"optional" json:"url" yaml:"url"`
	// user_bucket block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/ebs_snapshot_import#user_bucket EbsSnapshotImport#user_bucket}
	UserBucket *EbsSnapshotImportDiskContainerUserBucket `field:"optional" json:"userBucket" yaml:"userBucket"`
}

