// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package backupplan


type BackupPlanRule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.39.0/docs/resources/backup_plan#rule_name BackupPlan#rule_name}.
	RuleName *string `field:"required" json:"ruleName" yaml:"ruleName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.39.0/docs/resources/backup_plan#target_vault_name BackupPlan#target_vault_name}.
	TargetVaultName *string `field:"required" json:"targetVaultName" yaml:"targetVaultName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.39.0/docs/resources/backup_plan#completion_window BackupPlan#completion_window}.
	CompletionWindow *float64 `field:"optional" json:"completionWindow" yaml:"completionWindow"`
	// copy_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.39.0/docs/resources/backup_plan#copy_action BackupPlan#copy_action}
	CopyAction interface{} `field:"optional" json:"copyAction" yaml:"copyAction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.39.0/docs/resources/backup_plan#enable_continuous_backup BackupPlan#enable_continuous_backup}.
	EnableContinuousBackup interface{} `field:"optional" json:"enableContinuousBackup" yaml:"enableContinuousBackup"`
	// lifecycle block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.39.0/docs/resources/backup_plan#lifecycle BackupPlan#lifecycle}
	Lifecycle *BackupPlanRuleLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.39.0/docs/resources/backup_plan#recovery_point_tags BackupPlan#recovery_point_tags}.
	RecoveryPointTags *map[string]*string `field:"optional" json:"recoveryPointTags" yaml:"recoveryPointTags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.39.0/docs/resources/backup_plan#schedule BackupPlan#schedule}.
	Schedule *string `field:"optional" json:"schedule" yaml:"schedule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.39.0/docs/resources/backup_plan#start_window BackupPlan#start_window}.
	StartWindow *float64 `field:"optional" json:"startWindow" yaml:"startWindow"`
}

