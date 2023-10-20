// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kendraquerysuggestionsblocklist


type KendraQuerySuggestionsBlockListTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.22.0/docs/resources/kendra_query_suggestions_block_list#create KendraQuerySuggestionsBlockList#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.22.0/docs/resources/kendra_query_suggestions_block_list#delete KendraQuerySuggestionsBlockList#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.22.0/docs/resources/kendra_query_suggestions_block_list#update KendraQuerySuggestionsBlockList#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

