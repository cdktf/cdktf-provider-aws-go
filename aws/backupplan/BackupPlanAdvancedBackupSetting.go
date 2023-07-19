package backupplan


type BackupPlanAdvancedBackupSetting struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.8.0/docs/resources/backup_plan#backup_options BackupPlan#backup_options}.
	BackupOptions *map[string]*string `field:"required" json:"backupOptions" yaml:"backupOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.8.0/docs/resources/backup_plan#resource_type BackupPlan#resource_type}.
	ResourceType *string `field:"required" json:"resourceType" yaml:"resourceType"`
}

