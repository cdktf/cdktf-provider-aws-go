// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider


type AwsProviderDefaultTags struct {
	// Resource tags to default across all resources. Can also be configured with environment variables like `TF_AWS_DEFAULT_TAGS_<tag_name>`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.1/docs#tags AwsProvider#tags}
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

