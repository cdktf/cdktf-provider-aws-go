// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakerdomain


type SagemakerDomainDefaultUserSettingsCustomFileSystemConfig struct {
	// efs_file_system_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.98.0/docs/resources/sagemaker_domain#efs_file_system_config SagemakerDomain#efs_file_system_config}
	EfsFileSystemConfig *SagemakerDomainDefaultUserSettingsCustomFileSystemConfigEfsFileSystemConfig `field:"optional" json:"efsFileSystemConfig" yaml:"efsFileSystemConfig"`
}

