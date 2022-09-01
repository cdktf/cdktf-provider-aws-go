package backup


type BackupFrameworkControl struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/backup_framework#name BackupFramework#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// input_parameter block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/backup_framework#input_parameter BackupFramework#input_parameter}
	InputParameter interface{} `field:"optional" json:"inputParameter" yaml:"inputParameter"`
	// scope block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/backup_framework#scope BackupFramework#scope}
	Scope *BackupFrameworkControlScope `field:"optional" json:"scope" yaml:"scope"`
}

