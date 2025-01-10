// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakerdomain


type SagemakerDomainDefaultSpaceSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.0/docs/resources/sagemaker_domain#execution_role SagemakerDomain#execution_role}.
	ExecutionRole *string `field:"required" json:"executionRole" yaml:"executionRole"`
	// custom_file_system_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.0/docs/resources/sagemaker_domain#custom_file_system_config SagemakerDomain#custom_file_system_config}
	CustomFileSystemConfig interface{} `field:"optional" json:"customFileSystemConfig" yaml:"customFileSystemConfig"`
	// custom_posix_user_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.0/docs/resources/sagemaker_domain#custom_posix_user_config SagemakerDomain#custom_posix_user_config}
	CustomPosixUserConfig *SagemakerDomainDefaultSpaceSettingsCustomPosixUserConfig `field:"optional" json:"customPosixUserConfig" yaml:"customPosixUserConfig"`
	// jupyter_lab_app_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.0/docs/resources/sagemaker_domain#jupyter_lab_app_settings SagemakerDomain#jupyter_lab_app_settings}
	JupyterLabAppSettings *SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettings `field:"optional" json:"jupyterLabAppSettings" yaml:"jupyterLabAppSettings"`
	// jupyter_server_app_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.0/docs/resources/sagemaker_domain#jupyter_server_app_settings SagemakerDomain#jupyter_server_app_settings}
	JupyterServerAppSettings *SagemakerDomainDefaultSpaceSettingsJupyterServerAppSettings `field:"optional" json:"jupyterServerAppSettings" yaml:"jupyterServerAppSettings"`
	// kernel_gateway_app_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.0/docs/resources/sagemaker_domain#kernel_gateway_app_settings SagemakerDomain#kernel_gateway_app_settings}
	KernelGatewayAppSettings *SagemakerDomainDefaultSpaceSettingsKernelGatewayAppSettings `field:"optional" json:"kernelGatewayAppSettings" yaml:"kernelGatewayAppSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.0/docs/resources/sagemaker_domain#security_groups SagemakerDomain#security_groups}.
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// space_storage_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.0/docs/resources/sagemaker_domain#space_storage_settings SagemakerDomain#space_storage_settings}
	SpaceStorageSettings *SagemakerDomainDefaultSpaceSettingsSpaceStorageSettings `field:"optional" json:"spaceStorageSettings" yaml:"spaceStorageSettings"`
}

