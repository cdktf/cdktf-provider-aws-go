package gluedatacatalogencryptionsettings


type GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettings struct {
	// connection_password_encryption block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_data_catalog_encryption_settings#connection_password_encryption GlueDataCatalogEncryptionSettings#connection_password_encryption}
	ConnectionPasswordEncryption *GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsConnectionPasswordEncryption `field:"required" json:"connectionPasswordEncryption" yaml:"connectionPasswordEncryption"`
	// encryption_at_rest block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_data_catalog_encryption_settings#encryption_at_rest GlueDataCatalogEncryptionSettings#encryption_at_rest}
	EncryptionAtRest *GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsEncryptionAtRest `field:"required" json:"encryptionAtRest" yaml:"encryptionAtRest"`
}

