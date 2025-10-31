// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package backupplan


type BackupPlanRuleCopyActionLifecycle struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/backup_plan#cold_storage_after BackupPlan#cold_storage_after}.
	ColdStorageAfter *float64 `field:"optional" json:"coldStorageAfter" yaml:"coldStorageAfter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/backup_plan#delete_after BackupPlan#delete_after}.
	DeleteAfter *float64 `field:"optional" json:"deleteAfter" yaml:"deleteAfter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/backup_plan#opt_in_to_archive_for_supported_resources BackupPlan#opt_in_to_archive_for_supported_resources}.
	OptInToArchiveForSupportedResources interface{} `field:"optional" json:"optInToArchiveForSupportedResources" yaml:"optInToArchiveForSupportedResources"`
}

