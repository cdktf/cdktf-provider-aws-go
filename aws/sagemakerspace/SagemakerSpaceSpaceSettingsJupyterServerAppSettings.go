// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakerspace


type SagemakerSpaceSpaceSettingsJupyterServerAppSettings struct {
	// default_resource_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/sagemaker_space#default_resource_spec SagemakerSpace#default_resource_spec}
	DefaultResourceSpec *SagemakerSpaceSpaceSettingsJupyterServerAppSettingsDefaultResourceSpec `field:"required" json:"defaultResourceSpec" yaml:"defaultResourceSpec"`
	// code_repository block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/sagemaker_space#code_repository SagemakerSpace#code_repository}
	CodeRepository interface{} `field:"optional" json:"codeRepository" yaml:"codeRepository"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/sagemaker_space#lifecycle_config_arns SagemakerSpace#lifecycle_config_arns}.
	LifecycleConfigArns *[]*string `field:"optional" json:"lifecycleConfigArns" yaml:"lifecycleConfigArns"`
}

