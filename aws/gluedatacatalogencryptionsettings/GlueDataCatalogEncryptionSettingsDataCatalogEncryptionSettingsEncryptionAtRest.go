package gluedatacatalogencryptionsettings


type GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsEncryptionAtRest struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_data_catalog_encryption_settings#catalog_encryption_mode GlueDataCatalogEncryptionSettings#catalog_encryption_mode}.
	CatalogEncryptionMode *string `field:"required" json:"catalogEncryptionMode" yaml:"catalogEncryptionMode"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_data_catalog_encryption_settings#sse_aws_kms_key_id GlueDataCatalogEncryptionSettings#sse_aws_kms_key_id}.
	SseAwsKmsKeyId *string `field:"optional" json:"sseAwsKmsKeyId" yaml:"sseAwsKmsKeyId"`
}

