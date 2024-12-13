// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakerendpoint


type SagemakerEndpointDeploymentConfigBlueGreenUpdatePolicyTrafficRoutingConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.81.0/docs/resources/sagemaker_endpoint#type SagemakerEndpoint#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.81.0/docs/resources/sagemaker_endpoint#wait_interval_in_seconds SagemakerEndpoint#wait_interval_in_seconds}.
	WaitIntervalInSeconds *float64 `field:"required" json:"waitIntervalInSeconds" yaml:"waitIntervalInSeconds"`
	// canary_size block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.81.0/docs/resources/sagemaker_endpoint#canary_size SagemakerEndpoint#canary_size}
	CanarySize *SagemakerEndpointDeploymentConfigBlueGreenUpdatePolicyTrafficRoutingConfigurationCanarySize `field:"optional" json:"canarySize" yaml:"canarySize"`
	// linear_step_size block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.81.0/docs/resources/sagemaker_endpoint#linear_step_size SagemakerEndpoint#linear_step_size}
	LinearStepSize *SagemakerEndpointDeploymentConfigBlueGreenUpdatePolicyTrafficRoutingConfigurationLinearStepSize `field:"optional" json:"linearStepSize" yaml:"linearStepSize"`
}

