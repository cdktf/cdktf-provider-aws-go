// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package qbusinessapplication


type QbusinessApplicationEncryptionConfiguration struct {
	// The identifier of the AWS KMS key that is used to encrypt your data.
	//
	// Amazon Q doesn't support asymmetric keys.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.88.0/docs/resources/qbusiness_application#kms_key_id QbusinessApplication#kms_key_id}
	KmsKeyId *string `field:"required" json:"kmsKeyId" yaml:"kmsKeyId"`
}

