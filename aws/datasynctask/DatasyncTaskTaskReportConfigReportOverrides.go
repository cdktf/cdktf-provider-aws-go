// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasynctask


type DatasyncTaskTaskReportConfigReportOverrides struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.44.0/docs/resources/datasync_task#deleted_override DatasyncTask#deleted_override}.
	DeletedOverride *string `field:"optional" json:"deletedOverride" yaml:"deletedOverride"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.44.0/docs/resources/datasync_task#skipped_override DatasyncTask#skipped_override}.
	SkippedOverride *string `field:"optional" json:"skippedOverride" yaml:"skippedOverride"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.44.0/docs/resources/datasync_task#transferred_override DatasyncTask#transferred_override}.
	TransferredOverride *string `field:"optional" json:"transferredOverride" yaml:"transferredOverride"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.44.0/docs/resources/datasync_task#verified_override DatasyncTask#verified_override}.
	VerifiedOverride *string `field:"optional" json:"verifiedOverride" yaml:"verifiedOverride"`
}

