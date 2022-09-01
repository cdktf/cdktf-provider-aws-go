package sagemaker


type SagemakerEndpointConfigurationAsyncInferenceConfig struct {
	// output_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/sagemaker_endpoint_configuration#output_config SagemakerEndpointConfiguration#output_config}
	OutputConfig *SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfig `field:"required" json:"outputConfig" yaml:"outputConfig"`
	// client_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/sagemaker_endpoint_configuration#client_config SagemakerEndpointConfiguration#client_config}
	ClientConfig *SagemakerEndpointConfigurationAsyncInferenceConfigClientConfig `field:"optional" json:"clientConfig" yaml:"clientConfig"`
}

