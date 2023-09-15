// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakerendpointconfiguration


type SagemakerEndpointConfigurationAsyncInferenceConfig struct {
	// output_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.17.0/docs/resources/sagemaker_endpoint_configuration#output_config SagemakerEndpointConfiguration#output_config}
	OutputConfig *SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfig `field:"required" json:"outputConfig" yaml:"outputConfig"`
	// client_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.17.0/docs/resources/sagemaker_endpoint_configuration#client_config SagemakerEndpointConfiguration#client_config}
	ClientConfig *SagemakerEndpointConfigurationAsyncInferenceConfigClientConfig `field:"optional" json:"clientConfig" yaml:"clientConfig"`
}

