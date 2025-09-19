// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package backuprestoretestingplan


type BackupRestoreTestingPlanRecoveryPointSelection struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/backup_restore_testing_plan#algorithm BackupRestoreTestingPlan#algorithm}.
	Algorithm *string `field:"required" json:"algorithm" yaml:"algorithm"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/backup_restore_testing_plan#include_vaults BackupRestoreTestingPlan#include_vaults}.
	IncludeVaults *[]*string `field:"required" json:"includeVaults" yaml:"includeVaults"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/backup_restore_testing_plan#recovery_point_types BackupRestoreTestingPlan#recovery_point_types}.
	RecoveryPointTypes *[]*string `field:"required" json:"recoveryPointTypes" yaml:"recoveryPointTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/backup_restore_testing_plan#exclude_vaults BackupRestoreTestingPlan#exclude_vaults}.
	ExcludeVaults *[]*string `field:"optional" json:"excludeVaults" yaml:"excludeVaults"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/backup_restore_testing_plan#selection_window_days BackupRestoreTestingPlan#selection_window_days}.
	SelectionWindowDays *float64 `field:"optional" json:"selectionWindowDays" yaml:"selectionWindowDays"`
}

