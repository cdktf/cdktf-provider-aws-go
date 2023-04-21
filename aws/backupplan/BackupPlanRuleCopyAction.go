package backupplan


type BackupPlanRuleCopyAction struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.64.0/docs/resources/backup_plan#destination_vault_arn BackupPlan#destination_vault_arn}.
	DestinationVaultArn *string `field:"required" json:"destinationVaultArn" yaml:"destinationVaultArn"`
	// lifecycle block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.64.0/docs/resources/backup_plan#lifecycle BackupPlan#lifecycle}
	Lifecycle *BackupPlanRuleCopyActionLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
}

