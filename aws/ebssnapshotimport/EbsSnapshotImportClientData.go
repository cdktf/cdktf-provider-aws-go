// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ebssnapshotimport


type EbsSnapshotImportClientData struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/ebs_snapshot_import#comment EbsSnapshotImport#comment}.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/ebs_snapshot_import#upload_end EbsSnapshotImport#upload_end}.
	UploadEnd *string `field:"optional" json:"uploadEnd" yaml:"uploadEnd"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/ebs_snapshot_import#upload_size EbsSnapshotImport#upload_size}.
	UploadSize *float64 `field:"optional" json:"uploadSize" yaml:"uploadSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/ebs_snapshot_import#upload_start EbsSnapshotImport#upload_start}.
	UploadStart *string `field:"optional" json:"uploadStart" yaml:"uploadStart"`
}

