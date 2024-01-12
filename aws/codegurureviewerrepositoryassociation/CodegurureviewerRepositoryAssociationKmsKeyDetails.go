// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package codegurureviewerrepositoryassociation


type CodegurureviewerRepositoryAssociationKmsKeyDetails struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.0/docs/resources/codegurureviewer_repository_association#encryption_option CodegurureviewerRepositoryAssociation#encryption_option}.
	EncryptionOption *string `field:"optional" json:"encryptionOption" yaml:"encryptionOption"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.0/docs/resources/codegurureviewer_repository_association#kms_key_id CodegurureviewerRepositoryAssociation#kms_key_id}.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
}

