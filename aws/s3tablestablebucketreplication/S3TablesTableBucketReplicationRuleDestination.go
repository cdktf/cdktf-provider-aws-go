// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3tablestablebucketreplication


type S3TablesTableBucketReplicationRuleDestination struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/s3tables_table_bucket_replication#destination_table_bucket_arn S3TablesTableBucketReplication#destination_table_bucket_arn}.
	DestinationTableBucketArn *string `field:"required" json:"destinationTableBucketArn" yaml:"destinationTableBucketArn"`
}

