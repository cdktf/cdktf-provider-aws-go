package backupvault


type BackupVaultTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/backup_vault#delete BackupVault#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

