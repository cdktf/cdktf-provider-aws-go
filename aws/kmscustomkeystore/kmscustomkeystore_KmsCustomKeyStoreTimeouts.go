package kmscustomkeystore


type KmsCustomKeyStoreTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/kms_custom_key_store#create KmsCustomKeyStore#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/kms_custom_key_store#delete KmsCustomKeyStore#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/kms_custom_key_store#update KmsCustomKeyStore#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

