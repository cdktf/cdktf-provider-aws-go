// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3bucketreplicationconfiguration


type S3BucketReplicationConfigurationRuleDestinationMetrics struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/s3_bucket_replication_configuration#status S3BucketReplicationConfigurationA#status}.
	Status *string `field:"required" json:"status" yaml:"status"`
	// event_threshold block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/s3_bucket_replication_configuration#event_threshold S3BucketReplicationConfigurationA#event_threshold}
	EventThreshold *S3BucketReplicationConfigurationRuleDestinationMetricsEventThreshold `field:"optional" json:"eventThreshold" yaml:"eventThreshold"`
}

