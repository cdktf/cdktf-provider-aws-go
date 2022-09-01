package backup


type BackupSelectionConditionStringNotLike struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/backup_selection#key BackupSelection#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/backup_selection#value BackupSelection#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

