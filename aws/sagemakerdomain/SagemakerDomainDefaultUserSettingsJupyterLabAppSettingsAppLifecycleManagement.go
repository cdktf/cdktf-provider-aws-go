// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakerdomain


type SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsAppLifecycleManagement struct {
	// idle_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.97.0/docs/resources/sagemaker_domain#idle_settings SagemakerDomain#idle_settings}
	IdleSettings *SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsAppLifecycleManagementIdleSettings `field:"optional" json:"idleSettings" yaml:"idleSettings"`
}

