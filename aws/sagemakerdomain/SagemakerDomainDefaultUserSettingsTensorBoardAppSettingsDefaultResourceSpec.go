// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakerdomain


type SagemakerDomainDefaultUserSettingsTensorBoardAppSettingsDefaultResourceSpec struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.27.0/docs/resources/sagemaker_domain#instance_type SagemakerDomain#instance_type}.
	InstanceType *string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.27.0/docs/resources/sagemaker_domain#lifecycle_config_arn SagemakerDomain#lifecycle_config_arn}.
	LifecycleConfigArn *string `field:"optional" json:"lifecycleConfigArn" yaml:"lifecycleConfigArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.27.0/docs/resources/sagemaker_domain#sagemaker_image_arn SagemakerDomain#sagemaker_image_arn}.
	SagemakerImageArn *string `field:"optional" json:"sagemakerImageArn" yaml:"sagemakerImageArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.27.0/docs/resources/sagemaker_domain#sagemaker_image_version_arn SagemakerDomain#sagemaker_image_version_arn}.
	SagemakerImageVersionArn *string `field:"optional" json:"sagemakerImageVersionArn" yaml:"sagemakerImageVersionArn"`
}

