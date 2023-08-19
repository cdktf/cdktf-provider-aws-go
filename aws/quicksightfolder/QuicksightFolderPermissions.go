package quicksightfolder


type QuicksightFolderPermissions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/quicksight_folder#actions QuicksightFolder#actions}.
	Actions *[]*string `field:"required" json:"actions" yaml:"actions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/quicksight_folder#principal QuicksightFolder#principal}.
	Principal *string `field:"required" json:"principal" yaml:"principal"`
}

