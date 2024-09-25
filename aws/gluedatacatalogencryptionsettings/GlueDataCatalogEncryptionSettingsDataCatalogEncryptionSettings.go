// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gluedatacatalogencryptionsettings


type GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettings struct {
	// connection_password_encryption block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.68.0/docs/resources/glue_data_catalog_encryption_settings#connection_password_encryption GlueDataCatalogEncryptionSettings#connection_password_encryption}
	ConnectionPasswordEncryption *GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsConnectionPasswordEncryption `field:"required" json:"connectionPasswordEncryption" yaml:"connectionPasswordEncryption"`
	// encryption_at_rest block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.68.0/docs/resources/glue_data_catalog_encryption_settings#encryption_at_rest GlueDataCatalogEncryptionSettings#encryption_at_rest}
	EncryptionAtRest *GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsEncryptionAtRest `field:"required" json:"encryptionAtRest" yaml:"encryptionAtRest"`
}

