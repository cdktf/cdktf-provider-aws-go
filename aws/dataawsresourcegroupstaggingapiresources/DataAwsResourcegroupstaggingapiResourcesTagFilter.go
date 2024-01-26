// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsresourcegroupstaggingapiresources


type DataAwsResourcegroupstaggingapiResourcesTagFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.34.0/docs/data-sources/resourcegroupstaggingapi_resources#key DataAwsResourcegroupstaggingapiResources#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.34.0/docs/data-sources/resourcegroupstaggingapi_resources#values DataAwsResourcegroupstaggingapiResources#values}.
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

