package imagebuilder


type ImagebuilderImagePipelineSchedule struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/imagebuilder_image_pipeline#schedule_expression ImagebuilderImagePipeline#schedule_expression}.
	ScheduleExpression *string `field:"required" json:"scheduleExpression" yaml:"scheduleExpression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/imagebuilder_image_pipeline#pipeline_execution_start_condition ImagebuilderImagePipeline#pipeline_execution_start_condition}.
	PipelineExecutionStartCondition *string `field:"optional" json:"pipelineExecutionStartCondition" yaml:"pipelineExecutionStartCondition"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/imagebuilder_image_pipeline#timezone ImagebuilderImagePipeline#timezone}.
	Timezone *string `field:"optional" json:"timezone" yaml:"timezone"`
}

