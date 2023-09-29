// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rdscluster


type RdsClusterS3Import struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.19.0/docs/resources/rds_cluster#bucket_name RdsCluster#bucket_name}.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.19.0/docs/resources/rds_cluster#ingestion_role RdsCluster#ingestion_role}.
	IngestionRole *string `field:"required" json:"ingestionRole" yaml:"ingestionRole"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.19.0/docs/resources/rds_cluster#source_engine RdsCluster#source_engine}.
	SourceEngine *string `field:"required" json:"sourceEngine" yaml:"sourceEngine"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.19.0/docs/resources/rds_cluster#source_engine_version RdsCluster#source_engine_version}.
	SourceEngineVersion *string `field:"required" json:"sourceEngineVersion" yaml:"sourceEngineVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.19.0/docs/resources/rds_cluster#bucket_prefix RdsCluster#bucket_prefix}.
	BucketPrefix *string `field:"optional" json:"bucketPrefix" yaml:"bucketPrefix"`
}

