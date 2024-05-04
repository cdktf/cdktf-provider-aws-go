// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package osispipeline


type OsisPipelineEncryptionAtRestOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.48.0/docs/resources/osis_pipeline#kms_key_arn OsisPipeline#kms_key_arn}.
	KmsKeyArn *string `field:"required" json:"kmsKeyArn" yaml:"kmsKeyArn"`
}

