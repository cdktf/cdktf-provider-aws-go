// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package chimesdkmediapipelinesmediainsightspipelineconfiguration


type ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#language_code ChimesdkmediapipelinesMediaInsightsPipelineConfiguration#language_code}.
	LanguageCode *string `field:"required" json:"languageCode" yaml:"languageCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#content_identification_type ChimesdkmediapipelinesMediaInsightsPipelineConfiguration#content_identification_type}.
	ContentIdentificationType *string `field:"optional" json:"contentIdentificationType" yaml:"contentIdentificationType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#content_redaction_type ChimesdkmediapipelinesMediaInsightsPipelineConfiguration#content_redaction_type}.
	ContentRedactionType *string `field:"optional" json:"contentRedactionType" yaml:"contentRedactionType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#enable_partial_results_stabilization ChimesdkmediapipelinesMediaInsightsPipelineConfiguration#enable_partial_results_stabilization}.
	EnablePartialResultsStabilization interface{} `field:"optional" json:"enablePartialResultsStabilization" yaml:"enablePartialResultsStabilization"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#filter_partial_results ChimesdkmediapipelinesMediaInsightsPipelineConfiguration#filter_partial_results}.
	FilterPartialResults interface{} `field:"optional" json:"filterPartialResults" yaml:"filterPartialResults"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#language_model_name ChimesdkmediapipelinesMediaInsightsPipelineConfiguration#language_model_name}.
	LanguageModelName *string `field:"optional" json:"languageModelName" yaml:"languageModelName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#partial_results_stability ChimesdkmediapipelinesMediaInsightsPipelineConfiguration#partial_results_stability}.
	PartialResultsStability *string `field:"optional" json:"partialResultsStability" yaml:"partialResultsStability"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#pii_entity_types ChimesdkmediapipelinesMediaInsightsPipelineConfiguration#pii_entity_types}.
	PiiEntityTypes *string `field:"optional" json:"piiEntityTypes" yaml:"piiEntityTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#show_speaker_label ChimesdkmediapipelinesMediaInsightsPipelineConfiguration#show_speaker_label}.
	ShowSpeakerLabel interface{} `field:"optional" json:"showSpeakerLabel" yaml:"showSpeakerLabel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#vocabulary_filter_method ChimesdkmediapipelinesMediaInsightsPipelineConfiguration#vocabulary_filter_method}.
	VocabularyFilterMethod *string `field:"optional" json:"vocabularyFilterMethod" yaml:"vocabularyFilterMethod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#vocabulary_filter_name ChimesdkmediapipelinesMediaInsightsPipelineConfiguration#vocabulary_filter_name}.
	VocabularyFilterName *string `field:"optional" json:"vocabularyFilterName" yaml:"vocabularyFilterName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#vocabulary_name ChimesdkmediapipelinesMediaInsightsPipelineConfiguration#vocabulary_name}.
	VocabularyName *string `field:"optional" json:"vocabularyName" yaml:"vocabularyName"`
}

