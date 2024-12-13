// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gluesecurityconfiguration


type GlueSecurityConfigurationEncryptionConfigurationCloudwatchEncryption struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.81.0/docs/resources/glue_security_configuration#cloudwatch_encryption_mode GlueSecurityConfiguration#cloudwatch_encryption_mode}.
	CloudwatchEncryptionMode *string `field:"optional" json:"cloudwatchEncryptionMode" yaml:"cloudwatchEncryptionMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.81.0/docs/resources/glue_security_configuration#kms_key_arn GlueSecurityConfiguration#kms_key_arn}.
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
}

