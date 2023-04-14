package sagemakerdataqualityjobdefinition


type SagemakerDataQualityJobDefinitionDataQualityJobInputBatchTransformInputDatasetFormat struct {
	// csv block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/sagemaker_data_quality_job_definition#csv SagemakerDataQualityJobDefinition#csv}
	Csv *SagemakerDataQualityJobDefinitionDataQualityJobInputBatchTransformInputDatasetFormatCsv `field:"optional" json:"csv" yaml:"csv"`
	// json block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/sagemaker_data_quality_job_definition#json SagemakerDataQualityJobDefinition#json}
	Json *SagemakerDataQualityJobDefinitionDataQualityJobInputBatchTransformInputDatasetFormatJson `field:"optional" json:"json" yaml:"json"`
}

