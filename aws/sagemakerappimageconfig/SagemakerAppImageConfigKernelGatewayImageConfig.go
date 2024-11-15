// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakerappimageconfig


type SagemakerAppImageConfigKernelGatewayImageConfig struct {
	// kernel_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.76.0/docs/resources/sagemaker_app_image_config#kernel_spec SagemakerAppImageConfig#kernel_spec}
	KernelSpec interface{} `field:"required" json:"kernelSpec" yaml:"kernelSpec"`
	// file_system_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.76.0/docs/resources/sagemaker_app_image_config#file_system_config SagemakerAppImageConfig#file_system_config}
	FileSystemConfig *SagemakerAppImageConfigKernelGatewayImageConfigFileSystemConfig `field:"optional" json:"fileSystemConfig" yaml:"fileSystemConfig"`
}

