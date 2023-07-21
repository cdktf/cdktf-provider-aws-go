package fsxbackup


type FsxBackupTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/fsx_backup#create FsxBackup#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/fsx_backup#delete FsxBackup#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

