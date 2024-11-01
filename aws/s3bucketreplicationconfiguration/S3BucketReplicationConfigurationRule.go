// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3bucketreplicationconfiguration


type S3BucketReplicationConfigurationRule struct {
	// destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.74.0/docs/resources/s3_bucket_replication_configuration#destination S3BucketReplicationConfigurationA#destination}
	Destination *S3BucketReplicationConfigurationRuleDestination `field:"required" json:"destination" yaml:"destination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.74.0/docs/resources/s3_bucket_replication_configuration#status S3BucketReplicationConfigurationA#status}.
	Status *string `field:"required" json:"status" yaml:"status"`
	// delete_marker_replication block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.74.0/docs/resources/s3_bucket_replication_configuration#delete_marker_replication S3BucketReplicationConfigurationA#delete_marker_replication}
	DeleteMarkerReplication *S3BucketReplicationConfigurationRuleDeleteMarkerReplication `field:"optional" json:"deleteMarkerReplication" yaml:"deleteMarkerReplication"`
	// existing_object_replication block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.74.0/docs/resources/s3_bucket_replication_configuration#existing_object_replication S3BucketReplicationConfigurationA#existing_object_replication}
	ExistingObjectReplication *S3BucketReplicationConfigurationRuleExistingObjectReplication `field:"optional" json:"existingObjectReplication" yaml:"existingObjectReplication"`
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.74.0/docs/resources/s3_bucket_replication_configuration#filter S3BucketReplicationConfigurationA#filter}
	Filter *S3BucketReplicationConfigurationRuleFilter `field:"optional" json:"filter" yaml:"filter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.74.0/docs/resources/s3_bucket_replication_configuration#id S3BucketReplicationConfigurationA#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.74.0/docs/resources/s3_bucket_replication_configuration#prefix S3BucketReplicationConfigurationA#prefix}.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.74.0/docs/resources/s3_bucket_replication_configuration#priority S3BucketReplicationConfigurationA#priority}.
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// source_selection_criteria block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.74.0/docs/resources/s3_bucket_replication_configuration#source_selection_criteria S3BucketReplicationConfigurationA#source_selection_criteria}
	SourceSelectionCriteria *S3BucketReplicationConfigurationRuleSourceSelectionCriteria `field:"optional" json:"sourceSelectionCriteria" yaml:"sourceSelectionCriteria"`
}

