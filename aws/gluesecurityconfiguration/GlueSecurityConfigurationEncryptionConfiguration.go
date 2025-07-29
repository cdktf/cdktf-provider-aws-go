// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gluesecurityconfiguration


type GlueSecurityConfigurationEncryptionConfiguration struct {
	// cloudwatch_encryption block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.6.0/docs/resources/glue_security_configuration#cloudwatch_encryption GlueSecurityConfiguration#cloudwatch_encryption}
	CloudwatchEncryption *GlueSecurityConfigurationEncryptionConfigurationCloudwatchEncryption `field:"required" json:"cloudwatchEncryption" yaml:"cloudwatchEncryption"`
	// job_bookmarks_encryption block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.6.0/docs/resources/glue_security_configuration#job_bookmarks_encryption GlueSecurityConfiguration#job_bookmarks_encryption}
	JobBookmarksEncryption *GlueSecurityConfigurationEncryptionConfigurationJobBookmarksEncryption `field:"required" json:"jobBookmarksEncryption" yaml:"jobBookmarksEncryption"`
	// s3_encryption block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.6.0/docs/resources/glue_security_configuration#s3_encryption GlueSecurityConfiguration#s3_encryption}
	S3Encryption *GlueSecurityConfigurationEncryptionConfigurationS3Encryption `field:"required" json:"s3Encryption" yaml:"s3Encryption"`
}

