// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakerendpointconfiguration


type SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigNotificationConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.17.0/docs/resources/sagemaker_endpoint_configuration#error_topic SagemakerEndpointConfiguration#error_topic}.
	ErrorTopic *string `field:"optional" json:"errorTopic" yaml:"errorTopic"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.17.0/docs/resources/sagemaker_endpoint_configuration#include_inference_response_in SagemakerEndpointConfiguration#include_inference_response_in}.
	IncludeInferenceResponseIn *[]*string `field:"optional" json:"includeInferenceResponseIn" yaml:"includeInferenceResponseIn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.17.0/docs/resources/sagemaker_endpoint_configuration#success_topic SagemakerEndpointConfiguration#success_topic}.
	SuccessTopic *string `field:"optional" json:"successTopic" yaml:"successTopic"`
}

