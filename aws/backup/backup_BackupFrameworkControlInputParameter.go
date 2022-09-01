package backup


type BackupFrameworkControlInputParameter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/backup_framework#name BackupFramework#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/backup_framework#value BackupFramework#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

