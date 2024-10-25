// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakerendpoint


type SagemakerEndpointDeploymentConfigRollingUpdatePolicy struct {
	// maximum_batch_size block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.73.0/docs/resources/sagemaker_endpoint#maximum_batch_size SagemakerEndpoint#maximum_batch_size}
	MaximumBatchSize *SagemakerEndpointDeploymentConfigRollingUpdatePolicyMaximumBatchSize `field:"required" json:"maximumBatchSize" yaml:"maximumBatchSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.73.0/docs/resources/sagemaker_endpoint#wait_interval_in_seconds SagemakerEndpoint#wait_interval_in_seconds}.
	WaitIntervalInSeconds *float64 `field:"required" json:"waitIntervalInSeconds" yaml:"waitIntervalInSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.73.0/docs/resources/sagemaker_endpoint#maximum_execution_timeout_in_seconds SagemakerEndpoint#maximum_execution_timeout_in_seconds}.
	MaximumExecutionTimeoutInSeconds *float64 `field:"optional" json:"maximumExecutionTimeoutInSeconds" yaml:"maximumExecutionTimeoutInSeconds"`
	// rollback_maximum_batch_size block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.73.0/docs/resources/sagemaker_endpoint#rollback_maximum_batch_size SagemakerEndpoint#rollback_maximum_batch_size}
	RollbackMaximumBatchSize *SagemakerEndpointDeploymentConfigRollingUpdatePolicyRollbackMaximumBatchSize `field:"optional" json:"rollbackMaximumBatchSize" yaml:"rollbackMaximumBatchSize"`
}

