// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gluedatacatalogencryptionsettings


type GlueDataCatalogEncryptionSettingsDataCatalogEncryptionSettingsEncryptionAtRest struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.27.0/docs/resources/glue_data_catalog_encryption_settings#catalog_encryption_mode GlueDataCatalogEncryptionSettings#catalog_encryption_mode}.
	CatalogEncryptionMode *string `field:"required" json:"catalogEncryptionMode" yaml:"catalogEncryptionMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.27.0/docs/resources/glue_data_catalog_encryption_settings#sse_aws_kms_key_id GlueDataCatalogEncryptionSettings#sse_aws_kms_key_id}.
	SseAwsKmsKeyId *string `field:"optional" json:"sseAwsKmsKeyId" yaml:"sseAwsKmsKeyId"`
}

