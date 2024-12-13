// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawscetags


type DataAwsCeTagsSortBy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.81.0/docs/data-sources/ce_tags#key DataAwsCeTags#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.81.0/docs/data-sources/ce_tags#sort_order DataAwsCeTags#sort_order}.
	SortOrder *string `field:"optional" json:"sortOrder" yaml:"sortOrder"`
}

