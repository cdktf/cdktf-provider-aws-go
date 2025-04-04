// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakerdomain


type SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsAppLifecycleManagement struct {
	// idle_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.94.0/docs/resources/sagemaker_domain#idle_settings SagemakerDomain#idle_settings}
	IdleSettings *SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsAppLifecycleManagementIdleSettings `field:"optional" json:"idleSettings" yaml:"idleSettings"`
}

