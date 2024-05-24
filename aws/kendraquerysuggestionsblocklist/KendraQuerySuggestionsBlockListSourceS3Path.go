// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kendraquerysuggestionsblocklist


type KendraQuerySuggestionsBlockListSourceS3Path struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.51.0/docs/resources/kendra_query_suggestions_block_list#bucket KendraQuerySuggestionsBlockList#bucket}.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.51.0/docs/resources/kendra_query_suggestions_block_list#key KendraQuerySuggestionsBlockList#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
}

