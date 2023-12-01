// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3bucketmetric


type S3BucketMetricFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.29.0/docs/resources/s3_bucket_metric#prefix S3BucketMetric#prefix}.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.29.0/docs/resources/s3_bucket_metric#tags S3BucketMetric#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

