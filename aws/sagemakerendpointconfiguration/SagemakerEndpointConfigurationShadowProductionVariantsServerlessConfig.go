package sagemakerendpointconfiguration


type SagemakerEndpointConfigurationShadowProductionVariantsServerlessConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.64.0/docs/resources/sagemaker_endpoint_configuration#max_concurrency SagemakerEndpointConfiguration#max_concurrency}.
	MaxConcurrency *float64 `field:"required" json:"maxConcurrency" yaml:"maxConcurrency"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.64.0/docs/resources/sagemaker_endpoint_configuration#memory_size_in_mb SagemakerEndpointConfiguration#memory_size_in_mb}.
	MemorySizeInMb *float64 `field:"required" json:"memorySizeInMb" yaml:"memorySizeInMb"`
}

