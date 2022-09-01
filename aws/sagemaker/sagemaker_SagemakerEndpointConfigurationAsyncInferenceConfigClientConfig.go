package sagemaker


type SagemakerEndpointConfigurationAsyncInferenceConfigClientConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/sagemaker_endpoint_configuration#max_concurrent_invocations_per_instance SagemakerEndpointConfiguration#max_concurrent_invocations_per_instance}.
	MaxConcurrentInvocationsPerInstance *float64 `field:"optional" json:"maxConcurrentInvocationsPerInstance" yaml:"maxConcurrentInvocationsPerInstance"`
}

