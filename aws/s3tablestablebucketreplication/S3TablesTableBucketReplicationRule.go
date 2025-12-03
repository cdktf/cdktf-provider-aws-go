// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3tablestablebucketreplication


type S3TablesTableBucketReplicationRule struct {
	// destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/resources/s3tables_table_bucket_replication#destination S3TablesTableBucketReplication#destination}
	Destination interface{} `field:"optional" json:"destination" yaml:"destination"`
}

