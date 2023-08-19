package backupvault


type BackupVaultTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/backup_vault#delete BackupVault#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

