// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakeruserprofile


type SagemakerUserProfileUserSettingsJupyterLabAppSettingsAppLifecycleManagement struct {
	// idle_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.21.0/docs/resources/sagemaker_user_profile#idle_settings SagemakerUserProfile#idle_settings}
	IdleSettings *SagemakerUserProfileUserSettingsJupyterLabAppSettingsAppLifecycleManagementIdleSettings `field:"optional" json:"idleSettings" yaml:"idleSettings"`
}

