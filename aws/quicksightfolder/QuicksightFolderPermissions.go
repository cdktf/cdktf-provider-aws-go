package quicksightfolder


type QuicksightFolderPermissions struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_folder#actions QuicksightFolder#actions}.
	Actions *[]*string `field:"required" json:"actions" yaml:"actions"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_folder#principal QuicksightFolder#principal}.
	Principal *string `field:"required" json:"principal" yaml:"principal"`
}

