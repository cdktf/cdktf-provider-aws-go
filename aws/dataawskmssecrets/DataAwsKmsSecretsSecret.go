// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawskmssecrets


type DataAwsKmsSecretsSecret struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/data-sources/kms_secrets#name DataAwsKmsSecrets#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/data-sources/kms_secrets#payload DataAwsKmsSecrets#payload}.
	Payload *string `field:"required" json:"payload" yaml:"payload"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/data-sources/kms_secrets#context DataAwsKmsSecrets#context}.
	Context *map[string]*string `field:"optional" json:"context" yaml:"context"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/data-sources/kms_secrets#encryption_algorithm DataAwsKmsSecrets#encryption_algorithm}.
	EncryptionAlgorithm *string `field:"optional" json:"encryptionAlgorithm" yaml:"encryptionAlgorithm"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/data-sources/kms_secrets#grant_tokens DataAwsKmsSecrets#grant_tokens}.
	GrantTokens *[]*string `field:"optional" json:"grantTokens" yaml:"grantTokens"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/data-sources/kms_secrets#key_id DataAwsKmsSecrets#key_id}.
	KeyId *string `field:"optional" json:"keyId" yaml:"keyId"`
}

