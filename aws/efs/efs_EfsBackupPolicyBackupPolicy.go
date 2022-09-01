package efs


type EfsBackupPolicyBackupPolicy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/efs_backup_policy#status EfsBackupPolicy#status}.
	Status *string `field:"required" json:"status" yaml:"status"`
}

