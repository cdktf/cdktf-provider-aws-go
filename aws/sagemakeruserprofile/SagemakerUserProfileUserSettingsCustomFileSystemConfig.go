// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakeruserprofile


type SagemakerUserProfileUserSettingsCustomFileSystemConfig struct {
	// efs_file_system_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.0/docs/resources/sagemaker_user_profile#efs_file_system_config SagemakerUserProfile#efs_file_system_config}
	EfsFileSystemConfig interface{} `field:"optional" json:"efsFileSystemConfig" yaml:"efsFileSystemConfig"`
}

