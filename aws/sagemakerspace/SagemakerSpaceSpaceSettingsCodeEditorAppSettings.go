// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakerspace


type SagemakerSpaceSpaceSettingsCodeEditorAppSettings struct {
	// default_resource_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/sagemaker_space#default_resource_spec SagemakerSpace#default_resource_spec}
	DefaultResourceSpec *SagemakerSpaceSpaceSettingsCodeEditorAppSettingsDefaultResourceSpec `field:"required" json:"defaultResourceSpec" yaml:"defaultResourceSpec"`
	// app_lifecycle_management block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/sagemaker_space#app_lifecycle_management SagemakerSpace#app_lifecycle_management}
	AppLifecycleManagement *SagemakerSpaceSpaceSettingsCodeEditorAppSettingsAppLifecycleManagement `field:"optional" json:"appLifecycleManagement" yaml:"appLifecycleManagement"`
}

