// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakerspace


type SagemakerSpaceSpaceSettingsJupyterLabAppSettings struct {
	// default_resource_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.97.0/docs/resources/sagemaker_space#default_resource_spec SagemakerSpace#default_resource_spec}
	DefaultResourceSpec *SagemakerSpaceSpaceSettingsJupyterLabAppSettingsDefaultResourceSpec `field:"required" json:"defaultResourceSpec" yaml:"defaultResourceSpec"`
	// app_lifecycle_management block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.97.0/docs/resources/sagemaker_space#app_lifecycle_management SagemakerSpace#app_lifecycle_management}
	AppLifecycleManagement *SagemakerSpaceSpaceSettingsJupyterLabAppSettingsAppLifecycleManagement `field:"optional" json:"appLifecycleManagement" yaml:"appLifecycleManagement"`
	// code_repository block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.97.0/docs/resources/sagemaker_space#code_repository SagemakerSpace#code_repository}
	CodeRepository interface{} `field:"optional" json:"codeRepository" yaml:"codeRepository"`
}

