// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakerappimageconfig


type SagemakerAppImageConfigCodeEditorAppImageConfig struct {
	// container_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/sagemaker_app_image_config#container_config SagemakerAppImageConfig#container_config}
	ContainerConfig *SagemakerAppImageConfigCodeEditorAppImageConfigContainerConfig `field:"optional" json:"containerConfig" yaml:"containerConfig"`
	// file_system_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/sagemaker_app_image_config#file_system_config SagemakerAppImageConfig#file_system_config}
	FileSystemConfig *SagemakerAppImageConfigCodeEditorAppImageConfigFileSystemConfig `field:"optional" json:"fileSystemConfig" yaml:"fileSystemConfig"`
}

