package fsxbackup


type FsxBackupTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_backup#create FsxBackup#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_backup#delete FsxBackup#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

