// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakeruserprofile


type SagemakerUserProfileUserSettingsCodeEditorAppSettings struct {
	// default_resource_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/sagemaker_user_profile#default_resource_spec SagemakerUserProfile#default_resource_spec}
	DefaultResourceSpec *SagemakerUserProfileUserSettingsCodeEditorAppSettingsDefaultResourceSpec `field:"optional" json:"defaultResourceSpec" yaml:"defaultResourceSpec"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/sagemaker_user_profile#lifecycle_config_arns SagemakerUserProfile#lifecycle_config_arns}.
	LifecycleConfigArns *[]*string `field:"optional" json:"lifecycleConfigArns" yaml:"lifecycleConfigArns"`
}

