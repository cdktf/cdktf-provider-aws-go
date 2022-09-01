// Prebuilt aws Provider for Terraform CDK (cdktf)
package aws


type KendraQuerySuggestionsBlockListSourceS3Path struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/kendra_query_suggestions_block_list#bucket KendraQuerySuggestionsBlockList#bucket}.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/kendra_query_suggestions_block_list#key KendraQuerySuggestionsBlockList#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
}

