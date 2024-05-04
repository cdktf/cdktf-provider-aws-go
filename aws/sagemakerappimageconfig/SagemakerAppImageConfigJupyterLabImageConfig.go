// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakerappimageconfig


type SagemakerAppImageConfigJupyterLabImageConfig struct {
	// container_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.48.0/docs/resources/sagemaker_app_image_config#container_config SagemakerAppImageConfig#container_config}
	ContainerConfig *SagemakerAppImageConfigJupyterLabImageConfigContainerConfig `field:"optional" json:"containerConfig" yaml:"containerConfig"`
	// file_system_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.48.0/docs/resources/sagemaker_app_image_config#file_system_config SagemakerAppImageConfig#file_system_config}
	FileSystemConfig *SagemakerAppImageConfigJupyterLabImageConfigFileSystemConfig `field:"optional" json:"fileSystemConfig" yaml:"fileSystemConfig"`
}

