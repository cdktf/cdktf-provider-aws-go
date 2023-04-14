package sagemakerdataqualityjobdefinition


type SagemakerDataQualityJobDefinitionDataQualityJobOutputConfigMonitoringOutputs struct {
	// s3_output block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/sagemaker_data_quality_job_definition#s3_output SagemakerDataQualityJobDefinition#s3_output}
	S3Output *SagemakerDataQualityJobDefinitionDataQualityJobOutputConfigMonitoringOutputsS3Output `field:"required" json:"s3Output" yaml:"s3Output"`
}

