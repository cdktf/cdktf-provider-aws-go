// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package backuprestoretestingselection


type BackupRestoreTestingSelectionProtectedResourceConditionsStringEquals struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.86.0/docs/resources/backup_restore_testing_selection#key BackupRestoreTestingSelection#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.86.0/docs/resources/backup_restore_testing_selection#value BackupRestoreTestingSelection#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

