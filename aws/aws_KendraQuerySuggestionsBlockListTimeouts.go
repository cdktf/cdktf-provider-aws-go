// Prebuilt aws Provider for Terraform CDK (cdktf)
package aws


type KendraQuerySuggestionsBlockListTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/kendra_query_suggestions_block_list#create KendraQuerySuggestionsBlockList#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/kendra_query_suggestions_block_list#delete KendraQuerySuggestionsBlockList#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/kendra_query_suggestions_block_list#update KendraQuerySuggestionsBlockList#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

