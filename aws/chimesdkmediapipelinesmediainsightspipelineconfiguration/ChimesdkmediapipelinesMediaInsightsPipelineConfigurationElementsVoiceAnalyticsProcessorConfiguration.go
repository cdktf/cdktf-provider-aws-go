package chimesdkmediapipelinesmediainsightspipelineconfiguration


type ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsVoiceAnalyticsProcessorConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.10.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#speaker_search_status ChimesdkmediapipelinesMediaInsightsPipelineConfiguration#speaker_search_status}.
	SpeakerSearchStatus *string `field:"required" json:"speakerSearchStatus" yaml:"speakerSearchStatus"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.10.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#voice_tone_analysis_status ChimesdkmediapipelinesMediaInsightsPipelineConfiguration#voice_tone_analysis_status}.
	VoiceToneAnalysisStatus *string `field:"required" json:"voiceToneAnalysisStatus" yaml:"voiceToneAnalysisStatus"`
}

