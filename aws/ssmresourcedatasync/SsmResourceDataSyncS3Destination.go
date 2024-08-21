// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ssmresourcedatasync


type SsmResourceDataSyncS3Destination struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.1/docs/resources/ssm_resource_data_sync#bucket_name SsmResourceDataSync#bucket_name}.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.1/docs/resources/ssm_resource_data_sync#region SsmResourceDataSync#region}.
	Region *string `field:"required" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.1/docs/resources/ssm_resource_data_sync#kms_key_arn SsmResourceDataSync#kms_key_arn}.
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.1/docs/resources/ssm_resource_data_sync#prefix SsmResourceDataSync#prefix}.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.1/docs/resources/ssm_resource_data_sync#sync_format SsmResourceDataSync#sync_format}.
	SyncFormat *string `field:"optional" json:"syncFormat" yaml:"syncFormat"`
}

