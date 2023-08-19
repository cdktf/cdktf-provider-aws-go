package codebuildwebhook


type CodebuildWebhookFilterGroupFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/codebuild_webhook#pattern CodebuildWebhook#pattern}.
	Pattern *string `field:"required" json:"pattern" yaml:"pattern"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/codebuild_webhook#type CodebuildWebhook#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/codebuild_webhook#exclude_matched_pattern CodebuildWebhook#exclude_matched_pattern}.
	ExcludeMatchedPattern interface{} `field:"optional" json:"excludeMatchedPattern" yaml:"excludeMatchedPattern"`
}

