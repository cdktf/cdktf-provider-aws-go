// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package connectinstancestorageconfig


type ConnectInstanceStorageConfigStorageConfigS3ConfigEncryptionConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/connect_instance_storage_config#encryption_type ConnectInstanceStorageConfig#encryption_type}.
	EncryptionType *string `field:"required" json:"encryptionType" yaml:"encryptionType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/connect_instance_storage_config#key_id ConnectInstanceStorageConfig#key_id}.
	KeyId *string `field:"required" json:"keyId" yaml:"keyId"`
}

