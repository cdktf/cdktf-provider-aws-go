// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gluesecurityconfiguration


type GlueSecurityConfigurationEncryptionConfigurationJobBookmarksEncryption struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/glue_security_configuration#job_bookmarks_encryption_mode GlueSecurityConfiguration#job_bookmarks_encryption_mode}.
	JobBookmarksEncryptionMode *string `field:"optional" json:"jobBookmarksEncryptionMode" yaml:"jobBookmarksEncryptionMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/glue_security_configuration#kms_key_arn GlueSecurityConfiguration#kms_key_arn}.
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
}

