package backupplan


type BackupPlanRuleCopyAction struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/backup_plan#destination_vault_arn BackupPlan#destination_vault_arn}.
	DestinationVaultArn *string `field:"required" json:"destinationVaultArn" yaml:"destinationVaultArn"`
	// lifecycle block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/backup_plan#lifecycle BackupPlan#lifecycle}
	Lifecycle *BackupPlanRuleCopyActionLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
}

