package chimesdkmediapipelinesmediainsightspipelineconfiguration


type ChimesdkmediapipelinesMediaInsightsPipelineConfigurationRealTimeAlertConfigurationRules struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.10.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#type ChimesdkmediapipelinesMediaInsightsPipelineConfiguration#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// issue_detection_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.10.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#issue_detection_configuration ChimesdkmediapipelinesMediaInsightsPipelineConfiguration#issue_detection_configuration}
	IssueDetectionConfiguration *ChimesdkmediapipelinesMediaInsightsPipelineConfigurationRealTimeAlertConfigurationRulesIssueDetectionConfiguration `field:"optional" json:"issueDetectionConfiguration" yaml:"issueDetectionConfiguration"`
	// keyword_match_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.10.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#keyword_match_configuration ChimesdkmediapipelinesMediaInsightsPipelineConfiguration#keyword_match_configuration}
	KeywordMatchConfiguration *ChimesdkmediapipelinesMediaInsightsPipelineConfigurationRealTimeAlertConfigurationRulesKeywordMatchConfiguration `field:"optional" json:"keywordMatchConfiguration" yaml:"keywordMatchConfiguration"`
	// sentiment_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.10.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#sentiment_configuration ChimesdkmediapipelinesMediaInsightsPipelineConfiguration#sentiment_configuration}
	SentimentConfiguration *ChimesdkmediapipelinesMediaInsightsPipelineConfigurationRealTimeAlertConfigurationRulesSentimentConfiguration `field:"optional" json:"sentimentConfiguration" yaml:"sentimentConfiguration"`
}

