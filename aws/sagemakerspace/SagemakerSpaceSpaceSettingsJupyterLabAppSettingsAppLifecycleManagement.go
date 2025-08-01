// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakerspace


type SagemakerSpaceSpaceSettingsJupyterLabAppSettingsAppLifecycleManagement struct {
	// idle_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/sagemaker_space#idle_settings SagemakerSpace#idle_settings}
	IdleSettings *SagemakerSpaceSpaceSettingsJupyterLabAppSettingsAppLifecycleManagementIdleSettings `field:"optional" json:"idleSettings" yaml:"idleSettings"`
}

