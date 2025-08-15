// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3object


type S3ObjectOverrideProviderDefaultTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/resources/s3_object#tags S3Object#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

