// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3controldirectorybucketaccesspointscope


type S3ControlDirectoryBucketAccessPointScopeScope struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/s3control_directory_bucket_access_point_scope#permissions S3ControlDirectoryBucketAccessPointScope#permissions}.
	Permissions *[]*string `field:"optional" json:"permissions" yaml:"permissions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/s3control_directory_bucket_access_point_scope#prefixes S3ControlDirectoryBucketAccessPointScope#prefixes}.
	Prefixes *[]*string `field:"optional" json:"prefixes" yaml:"prefixes"`
}

