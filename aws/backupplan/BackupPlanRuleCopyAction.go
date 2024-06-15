// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package backupplan


type BackupPlanRuleCopyAction struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.54.1/docs/resources/backup_plan#destination_vault_arn BackupPlan#destination_vault_arn}.
	DestinationVaultArn *string `field:"required" json:"destinationVaultArn" yaml:"destinationVaultArn"`
	// lifecycle block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.54.1/docs/resources/backup_plan#lifecycle BackupPlan#lifecycle}
	Lifecycle *BackupPlanRuleCopyActionLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
}

