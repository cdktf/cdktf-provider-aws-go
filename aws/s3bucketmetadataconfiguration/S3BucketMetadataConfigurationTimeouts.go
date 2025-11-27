// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3bucketmetadataconfiguration


type S3BucketMetadataConfigurationTimeouts struct {
	// A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/s3_bucket_metadata_configuration#create S3BucketMetadataConfiguration#create}
	Create *string `field:"optional" json:"create" yaml:"create"`
}

