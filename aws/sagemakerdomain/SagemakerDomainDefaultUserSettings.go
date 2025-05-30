// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakerdomain


type SagemakerDomainDefaultUserSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/resources/sagemaker_domain#execution_role SagemakerDomain#execution_role}.
	ExecutionRole *string `field:"required" json:"executionRole" yaml:"executionRole"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/resources/sagemaker_domain#auto_mount_home_efs SagemakerDomain#auto_mount_home_efs}.
	AutoMountHomeEfs *string `field:"optional" json:"autoMountHomeEfs" yaml:"autoMountHomeEfs"`
	// canvas_app_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/resources/sagemaker_domain#canvas_app_settings SagemakerDomain#canvas_app_settings}
	CanvasAppSettings *SagemakerDomainDefaultUserSettingsCanvasAppSettings `field:"optional" json:"canvasAppSettings" yaml:"canvasAppSettings"`
	// code_editor_app_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/resources/sagemaker_domain#code_editor_app_settings SagemakerDomain#code_editor_app_settings}
	CodeEditorAppSettings *SagemakerDomainDefaultUserSettingsCodeEditorAppSettings `field:"optional" json:"codeEditorAppSettings" yaml:"codeEditorAppSettings"`
	// custom_file_system_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/resources/sagemaker_domain#custom_file_system_config SagemakerDomain#custom_file_system_config}
	CustomFileSystemConfig interface{} `field:"optional" json:"customFileSystemConfig" yaml:"customFileSystemConfig"`
	// custom_posix_user_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/resources/sagemaker_domain#custom_posix_user_config SagemakerDomain#custom_posix_user_config}
	CustomPosixUserConfig *SagemakerDomainDefaultUserSettingsCustomPosixUserConfig `field:"optional" json:"customPosixUserConfig" yaml:"customPosixUserConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/resources/sagemaker_domain#default_landing_uri SagemakerDomain#default_landing_uri}.
	DefaultLandingUri *string `field:"optional" json:"defaultLandingUri" yaml:"defaultLandingUri"`
	// jupyter_lab_app_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/resources/sagemaker_domain#jupyter_lab_app_settings SagemakerDomain#jupyter_lab_app_settings}
	JupyterLabAppSettings *SagemakerDomainDefaultUserSettingsJupyterLabAppSettings `field:"optional" json:"jupyterLabAppSettings" yaml:"jupyterLabAppSettings"`
	// jupyter_server_app_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/resources/sagemaker_domain#jupyter_server_app_settings SagemakerDomain#jupyter_server_app_settings}
	JupyterServerAppSettings *SagemakerDomainDefaultUserSettingsJupyterServerAppSettings `field:"optional" json:"jupyterServerAppSettings" yaml:"jupyterServerAppSettings"`
	// kernel_gateway_app_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/resources/sagemaker_domain#kernel_gateway_app_settings SagemakerDomain#kernel_gateway_app_settings}
	KernelGatewayAppSettings *SagemakerDomainDefaultUserSettingsKernelGatewayAppSettings `field:"optional" json:"kernelGatewayAppSettings" yaml:"kernelGatewayAppSettings"`
	// r_session_app_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/resources/sagemaker_domain#r_session_app_settings SagemakerDomain#r_session_app_settings}
	RSessionAppSettings *SagemakerDomainDefaultUserSettingsRSessionAppSettings `field:"optional" json:"rSessionAppSettings" yaml:"rSessionAppSettings"`
	// r_studio_server_pro_app_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/resources/sagemaker_domain#r_studio_server_pro_app_settings SagemakerDomain#r_studio_server_pro_app_settings}
	RStudioServerProAppSettings *SagemakerDomainDefaultUserSettingsRStudioServerProAppSettings `field:"optional" json:"rStudioServerProAppSettings" yaml:"rStudioServerProAppSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/resources/sagemaker_domain#security_groups SagemakerDomain#security_groups}.
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// sharing_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/resources/sagemaker_domain#sharing_settings SagemakerDomain#sharing_settings}
	SharingSettings *SagemakerDomainDefaultUserSettingsSharingSettings `field:"optional" json:"sharingSettings" yaml:"sharingSettings"`
	// space_storage_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/resources/sagemaker_domain#space_storage_settings SagemakerDomain#space_storage_settings}
	SpaceStorageSettings *SagemakerDomainDefaultUserSettingsSpaceStorageSettings `field:"optional" json:"spaceStorageSettings" yaml:"spaceStorageSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/resources/sagemaker_domain#studio_web_portal SagemakerDomain#studio_web_portal}.
	StudioWebPortal *string `field:"optional" json:"studioWebPortal" yaml:"studioWebPortal"`
	// studio_web_portal_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/resources/sagemaker_domain#studio_web_portal_settings SagemakerDomain#studio_web_portal_settings}
	StudioWebPortalSettings *SagemakerDomainDefaultUserSettingsStudioWebPortalSettings `field:"optional" json:"studioWebPortalSettings" yaml:"studioWebPortalSettings"`
	// tensor_board_app_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/resources/sagemaker_domain#tensor_board_app_settings SagemakerDomain#tensor_board_app_settings}
	TensorBoardAppSettings *SagemakerDomainDefaultUserSettingsTensorBoardAppSettings `field:"optional" json:"tensorBoardAppSettings" yaml:"tensorBoardAppSettings"`
}

