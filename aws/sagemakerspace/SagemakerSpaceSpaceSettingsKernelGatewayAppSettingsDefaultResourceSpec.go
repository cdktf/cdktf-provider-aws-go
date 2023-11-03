// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakerspace


type SagemakerSpaceSpaceSettingsKernelGatewayAppSettingsDefaultResourceSpec struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.24.0/docs/resources/sagemaker_space#instance_type SagemakerSpace#instance_type}.
	InstanceType *string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.24.0/docs/resources/sagemaker_space#lifecycle_config_arn SagemakerSpace#lifecycle_config_arn}.
	LifecycleConfigArn *string `field:"optional" json:"lifecycleConfigArn" yaml:"lifecycleConfigArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.24.0/docs/resources/sagemaker_space#sagemaker_image_arn SagemakerSpace#sagemaker_image_arn}.
	SagemakerImageArn *string `field:"optional" json:"sagemakerImageArn" yaml:"sagemakerImageArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.24.0/docs/resources/sagemaker_space#sagemaker_image_version_arn SagemakerSpace#sagemaker_image_version_arn}.
	SagemakerImageVersionArn *string `field:"optional" json:"sagemakerImageVersionArn" yaml:"sagemakerImageVersionArn"`
}

