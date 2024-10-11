// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakerendpoint


type SagemakerEndpointDeploymentConfig struct {
	// auto_rollback_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/sagemaker_endpoint#auto_rollback_configuration SagemakerEndpoint#auto_rollback_configuration}
	AutoRollbackConfiguration *SagemakerEndpointDeploymentConfigAutoRollbackConfiguration `field:"optional" json:"autoRollbackConfiguration" yaml:"autoRollbackConfiguration"`
	// blue_green_update_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/sagemaker_endpoint#blue_green_update_policy SagemakerEndpoint#blue_green_update_policy}
	BlueGreenUpdatePolicy *SagemakerEndpointDeploymentConfigBlueGreenUpdatePolicy `field:"optional" json:"blueGreenUpdatePolicy" yaml:"blueGreenUpdatePolicy"`
	// rolling_update_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/sagemaker_endpoint#rolling_update_policy SagemakerEndpoint#rolling_update_policy}
	RollingUpdatePolicy *SagemakerEndpointDeploymentConfigRollingUpdatePolicy `field:"optional" json:"rollingUpdatePolicy" yaml:"rollingUpdatePolicy"`
}

