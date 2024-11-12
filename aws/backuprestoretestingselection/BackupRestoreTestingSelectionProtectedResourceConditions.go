// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package backuprestoretestingselection


type BackupRestoreTestingSelectionProtectedResourceConditions struct {
	// string_equals block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/backup_restore_testing_selection#string_equals BackupRestoreTestingSelection#string_equals}
	StringEquals interface{} `field:"optional" json:"stringEquals" yaml:"stringEquals"`
	// string_not_equals block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/backup_restore_testing_selection#string_not_equals BackupRestoreTestingSelection#string_not_equals}
	StringNotEquals interface{} `field:"optional" json:"stringNotEquals" yaml:"stringNotEquals"`
}

