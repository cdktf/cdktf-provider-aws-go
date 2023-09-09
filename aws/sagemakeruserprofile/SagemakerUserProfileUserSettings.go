// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakeruserprofile


type SagemakerUserProfileUserSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.1/docs/resources/sagemaker_user_profile#execution_role SagemakerUserProfile#execution_role}.
	ExecutionRole *string `field:"required" json:"executionRole" yaml:"executionRole"`
	// canvas_app_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.1/docs/resources/sagemaker_user_profile#canvas_app_settings SagemakerUserProfile#canvas_app_settings}
	CanvasAppSettings *SagemakerUserProfileUserSettingsCanvasAppSettings `field:"optional" json:"canvasAppSettings" yaml:"canvasAppSettings"`
	// jupyter_server_app_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.1/docs/resources/sagemaker_user_profile#jupyter_server_app_settings SagemakerUserProfile#jupyter_server_app_settings}
	JupyterServerAppSettings *SagemakerUserProfileUserSettingsJupyterServerAppSettings `field:"optional" json:"jupyterServerAppSettings" yaml:"jupyterServerAppSettings"`
	// kernel_gateway_app_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.1/docs/resources/sagemaker_user_profile#kernel_gateway_app_settings SagemakerUserProfile#kernel_gateway_app_settings}
	KernelGatewayAppSettings *SagemakerUserProfileUserSettingsKernelGatewayAppSettings `field:"optional" json:"kernelGatewayAppSettings" yaml:"kernelGatewayAppSettings"`
	// r_session_app_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.1/docs/resources/sagemaker_user_profile#r_session_app_settings SagemakerUserProfile#r_session_app_settings}
	RSessionAppSettings *SagemakerUserProfileUserSettingsRSessionAppSettings `field:"optional" json:"rSessionAppSettings" yaml:"rSessionAppSettings"`
	// r_studio_server_pro_app_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.1/docs/resources/sagemaker_user_profile#r_studio_server_pro_app_settings SagemakerUserProfile#r_studio_server_pro_app_settings}
	RStudioServerProAppSettings *SagemakerUserProfileUserSettingsRStudioServerProAppSettings `field:"optional" json:"rStudioServerProAppSettings" yaml:"rStudioServerProAppSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.1/docs/resources/sagemaker_user_profile#security_groups SagemakerUserProfile#security_groups}.
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// sharing_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.1/docs/resources/sagemaker_user_profile#sharing_settings SagemakerUserProfile#sharing_settings}
	SharingSettings *SagemakerUserProfileUserSettingsSharingSettings `field:"optional" json:"sharingSettings" yaml:"sharingSettings"`
	// tensor_board_app_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.1/docs/resources/sagemaker_user_profile#tensor_board_app_settings SagemakerUserProfile#tensor_board_app_settings}
	TensorBoardAppSettings *SagemakerUserProfileUserSettingsTensorBoardAppSettings `field:"optional" json:"tensorBoardAppSettings" yaml:"tensorBoardAppSettings"`
}

