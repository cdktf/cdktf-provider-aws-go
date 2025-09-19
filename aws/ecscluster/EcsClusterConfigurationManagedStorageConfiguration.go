// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ecscluster


type EcsClusterConfigurationManagedStorageConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/ecs_cluster#fargate_ephemeral_storage_kms_key_id EcsCluster#fargate_ephemeral_storage_kms_key_id}.
	FargateEphemeralStorageKmsKeyId *string `field:"optional" json:"fargateEphemeralStorageKmsKeyId" yaml:"fargateEphemeralStorageKmsKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/ecs_cluster#kms_key_id EcsCluster#kms_key_id}.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
}

