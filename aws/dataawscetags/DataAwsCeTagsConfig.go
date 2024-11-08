// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawscetags

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAwsCeTagsConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// time_period block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.0/docs/data-sources/ce_tags#time_period DataAwsCeTags#time_period}
	TimePeriod *DataAwsCeTagsTimePeriod `field:"required" json:"timePeriod" yaml:"timePeriod"`
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.0/docs/data-sources/ce_tags#filter DataAwsCeTags#filter}
	Filter *DataAwsCeTagsFilter `field:"optional" json:"filter" yaml:"filter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.0/docs/data-sources/ce_tags#id DataAwsCeTags#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.0/docs/data-sources/ce_tags#search_string DataAwsCeTags#search_string}.
	SearchString *string `field:"optional" json:"searchString" yaml:"searchString"`
	// sort_by block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.0/docs/data-sources/ce_tags#sort_by DataAwsCeTags#sort_by}
	SortBy interface{} `field:"optional" json:"sortBy" yaml:"sortBy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.0/docs/data-sources/ce_tags#tag_key DataAwsCeTags#tag_key}.
	TagKey *string `field:"optional" json:"tagKey" yaml:"tagKey"`
}

