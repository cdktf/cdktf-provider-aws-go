package codepipelinecustomactiontype


type CodepipelineCustomActionTypeOutputArtifactDetails struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codepipeline_custom_action_type#maximum_count CodepipelineCustomActionType#maximum_count}.
	MaximumCount *float64 `field:"required" json:"maximumCount" yaml:"maximumCount"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codepipeline_custom_action_type#minimum_count CodepipelineCustomActionType#minimum_count}.
	MinimumCount *float64 `field:"required" json:"minimumCount" yaml:"minimumCount"`
}

