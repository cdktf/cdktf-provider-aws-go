// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3bucketabac


type S3BucketAbacAbacStatus struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/s3_bucket_abac#status S3BucketAbac#status}.
	Status *string `field:"required" json:"status" yaml:"status"`
}

