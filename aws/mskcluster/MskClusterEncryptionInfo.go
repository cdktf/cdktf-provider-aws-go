// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mskcluster


type MskClusterEncryptionInfo struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.41.0/docs/resources/msk_cluster#encryption_at_rest_kms_key_arn MskCluster#encryption_at_rest_kms_key_arn}.
	EncryptionAtRestKmsKeyArn *string `field:"optional" json:"encryptionAtRestKmsKeyArn" yaml:"encryptionAtRestKmsKeyArn"`
	// encryption_in_transit block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.41.0/docs/resources/msk_cluster#encryption_in_transit MskCluster#encryption_in_transit}
	EncryptionInTransit *MskClusterEncryptionInfoEncryptionInTransit `field:"optional" json:"encryptionInTransit" yaml:"encryptionInTransit"`
}

