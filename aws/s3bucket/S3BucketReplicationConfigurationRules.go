// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3bucket


type S3BucketReplicationConfigurationRules struct {
	// destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.56.1/docs/resources/s3_bucket#destination S3Bucket#destination}
	Destination *S3BucketReplicationConfigurationRulesDestination `field:"required" json:"destination" yaml:"destination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.56.1/docs/resources/s3_bucket#status S3Bucket#status}.
	Status *string `field:"required" json:"status" yaml:"status"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.56.1/docs/resources/s3_bucket#delete_marker_replication_status S3Bucket#delete_marker_replication_status}.
	DeleteMarkerReplicationStatus *string `field:"optional" json:"deleteMarkerReplicationStatus" yaml:"deleteMarkerReplicationStatus"`
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.56.1/docs/resources/s3_bucket#filter S3Bucket#filter}
	Filter *S3BucketReplicationConfigurationRulesFilter `field:"optional" json:"filter" yaml:"filter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.56.1/docs/resources/s3_bucket#id S3Bucket#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.56.1/docs/resources/s3_bucket#prefix S3Bucket#prefix}.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.56.1/docs/resources/s3_bucket#priority S3Bucket#priority}.
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// source_selection_criteria block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.56.1/docs/resources/s3_bucket#source_selection_criteria S3Bucket#source_selection_criteria}
	SourceSelectionCriteria *S3BucketReplicationConfigurationRulesSourceSelectionCriteria `field:"optional" json:"sourceSelectionCriteria" yaml:"sourceSelectionCriteria"`
}

