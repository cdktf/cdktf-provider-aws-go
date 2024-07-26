// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gluedatacatalogencryptionsettings


type GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsConnectionPasswordEncryption struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/glue_data_catalog_encryption_settings#return_connection_password_encrypted GlueDataCatalogEncryptionSettings#return_connection_password_encrypted}.
	ReturnConnectionPasswordEncrypted interface{} `field:"required" json:"returnConnectionPasswordEncrypted" yaml:"returnConnectionPasswordEncrypted"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/glue_data_catalog_encryption_settings#aws_kms_key_id GlueDataCatalogEncryptionSettings#aws_kms_key_id}.
	AwsKmsKeyId *string `field:"optional" json:"awsKmsKeyId" yaml:"awsKmsKeyId"`
}

