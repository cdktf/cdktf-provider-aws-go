package efsbackuppolicy


type EfsBackupPolicyBackupPolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/efs_backup_policy#status EfsBackupPolicy#status}.
	Status *string `field:"required" json:"status" yaml:"status"`
}

