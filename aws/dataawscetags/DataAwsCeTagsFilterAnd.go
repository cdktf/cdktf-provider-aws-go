// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawscetags


type DataAwsCeTagsFilterAnd struct {
	// cost_category block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.43.0/docs/data-sources/ce_tags#cost_category DataAwsCeTags#cost_category}
	CostCategory *DataAwsCeTagsFilterAndCostCategory `field:"optional" json:"costCategory" yaml:"costCategory"`
	// dimension block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.43.0/docs/data-sources/ce_tags#dimension DataAwsCeTags#dimension}
	Dimension *DataAwsCeTagsFilterAndDimension `field:"optional" json:"dimension" yaml:"dimension"`
	// tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.43.0/docs/data-sources/ce_tags#tags DataAwsCeTags#tags}
	Tags *DataAwsCeTagsFilterAndTags `field:"optional" json:"tags" yaml:"tags"`
}

