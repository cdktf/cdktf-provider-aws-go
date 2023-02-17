package backupplan


type BackupPlanRuleCopyActionLifecycle struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/backup_plan#cold_storage_after BackupPlan#cold_storage_after}.
	ColdStorageAfter *float64 `field:"optional" json:"coldStorageAfter" yaml:"coldStorageAfter"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/backup_plan#delete_after BackupPlan#delete_after}.
	DeleteAfter *float64 `field:"optional" json:"deleteAfter" yaml:"deleteAfter"`
}

