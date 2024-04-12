// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package chimesdkmediapipelinesmediainsightspipelineconfiguration


type ChimesdkmediapipelinesMediaInsightsPipelineConfigurationRealTimeAlertConfigurationRulesSentimentConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#rule_name ChimesdkmediapipelinesMediaInsightsPipelineConfiguration#rule_name}.
	RuleName *string `field:"required" json:"ruleName" yaml:"ruleName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#sentiment_type ChimesdkmediapipelinesMediaInsightsPipelineConfiguration#sentiment_type}.
	SentimentType *string `field:"required" json:"sentimentType" yaml:"sentimentType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#time_period ChimesdkmediapipelinesMediaInsightsPipelineConfiguration#time_period}.
	TimePeriod *float64 `field:"required" json:"timePeriod" yaml:"timePeriod"`
}

