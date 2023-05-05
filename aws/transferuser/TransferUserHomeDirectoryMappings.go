package transferuser


type TransferUserHomeDirectoryMappings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/transfer_user#entry TransferUser#entry}.
	Entry *string `field:"required" json:"entry" yaml:"entry"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/transfer_user#target TransferUser#target}.
	Target *string `field:"required" json:"target" yaml:"target"`
}

