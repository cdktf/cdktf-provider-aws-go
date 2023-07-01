package backupplan


type BackupPlanRuleLifecycle struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/resources/backup_plan#cold_storage_after BackupPlan#cold_storage_after}.
	ColdStorageAfter *float64 `field:"optional" json:"coldStorageAfter" yaml:"coldStorageAfter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/resources/backup_plan#delete_after BackupPlan#delete_after}.
	DeleteAfter *float64 `field:"optional" json:"deleteAfter" yaml:"deleteAfter"`
}

