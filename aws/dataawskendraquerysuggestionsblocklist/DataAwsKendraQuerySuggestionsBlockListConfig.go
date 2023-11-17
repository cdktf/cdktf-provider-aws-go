// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawskendraquerysuggestionsblocklist

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAwsKendraQuerySuggestionsBlockListConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.26.0/docs/data-sources/kendra_query_suggestions_block_list#index_id DataAwsKendraQuerySuggestionsBlockList#index_id}.
	IndexId *string `field:"required" json:"indexId" yaml:"indexId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.26.0/docs/data-sources/kendra_query_suggestions_block_list#query_suggestions_block_list_id DataAwsKendraQuerySuggestionsBlockList#query_suggestions_block_list_id}.
	QuerySuggestionsBlockListId *string `field:"required" json:"querySuggestionsBlockListId" yaml:"querySuggestionsBlockListId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.26.0/docs/data-sources/kendra_query_suggestions_block_list#id DataAwsKendraQuerySuggestionsBlockList#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.26.0/docs/data-sources/kendra_query_suggestions_block_list#tags DataAwsKendraQuerySuggestionsBlockList#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

