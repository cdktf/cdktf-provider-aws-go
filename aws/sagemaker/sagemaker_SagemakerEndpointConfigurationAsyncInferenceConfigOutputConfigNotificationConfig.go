package sagemaker


type SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigNotificationConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/sagemaker_endpoint_configuration#error_topic SagemakerEndpointConfiguration#error_topic}.
	ErrorTopic *string `field:"optional" json:"errorTopic" yaml:"errorTopic"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/sagemaker_endpoint_configuration#success_topic SagemakerEndpointConfiguration#success_topic}.
	SuccessTopic *string `field:"optional" json:"successTopic" yaml:"successTopic"`
}

