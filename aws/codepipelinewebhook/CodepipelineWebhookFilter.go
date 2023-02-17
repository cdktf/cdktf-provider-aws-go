package codepipelinewebhook


type CodepipelineWebhookFilter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codepipeline_webhook#json_path CodepipelineWebhook#json_path}.
	JsonPath *string `field:"required" json:"jsonPath" yaml:"jsonPath"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codepipeline_webhook#match_equals CodepipelineWebhook#match_equals}.
	MatchEquals *string `field:"required" json:"matchEquals" yaml:"matchEquals"`
}

