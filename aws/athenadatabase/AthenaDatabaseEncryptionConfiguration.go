package athenadatabase


type AthenaDatabaseEncryptionConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/athena_database#encryption_option AthenaDatabase#encryption_option}.
	EncryptionOption *string `field:"required" json:"encryptionOption" yaml:"encryptionOption"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/athena_database#kms_key AthenaDatabase#kms_key}.
	KmsKey *string `field:"optional" json:"kmsKey" yaml:"kmsKey"`
}

