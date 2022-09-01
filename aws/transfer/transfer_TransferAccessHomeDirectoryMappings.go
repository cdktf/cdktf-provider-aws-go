package transfer


type TransferAccessHomeDirectoryMappings struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/transfer_access#entry TransferAccess#entry}.
	Entry *string `field:"required" json:"entry" yaml:"entry"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/transfer_access#target TransferAccess#target}.
	Target *string `field:"required" json:"target" yaml:"target"`
}

