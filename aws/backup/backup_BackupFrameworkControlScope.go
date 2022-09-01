package backup


type BackupFrameworkControlScope struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/backup_framework#compliance_resource_ids BackupFramework#compliance_resource_ids}.
	ComplianceResourceIds *[]*string `field:"optional" json:"complianceResourceIds" yaml:"complianceResourceIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/backup_framework#compliance_resource_types BackupFramework#compliance_resource_types}.
	ComplianceResourceTypes *[]*string `field:"optional" json:"complianceResourceTypes" yaml:"complianceResourceTypes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/backup_framework#tags BackupFramework#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

