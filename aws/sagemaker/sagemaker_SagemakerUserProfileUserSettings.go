package sagemaker


type SagemakerUserProfileUserSettings struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/sagemaker_user_profile#execution_role SagemakerUserProfile#execution_role}.
	ExecutionRole *string `field:"required" json:"executionRole" yaml:"executionRole"`
	// jupyter_server_app_settings block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/sagemaker_user_profile#jupyter_server_app_settings SagemakerUserProfile#jupyter_server_app_settings}
	JupyterServerAppSettings *SagemakerUserProfileUserSettingsJupyterServerAppSettings `field:"optional" json:"jupyterServerAppSettings" yaml:"jupyterServerAppSettings"`
	// kernel_gateway_app_settings block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/sagemaker_user_profile#kernel_gateway_app_settings SagemakerUserProfile#kernel_gateway_app_settings}
	KernelGatewayAppSettings *SagemakerUserProfileUserSettingsKernelGatewayAppSettings `field:"optional" json:"kernelGatewayAppSettings" yaml:"kernelGatewayAppSettings"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/sagemaker_user_profile#security_groups SagemakerUserProfile#security_groups}.
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// sharing_settings block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/sagemaker_user_profile#sharing_settings SagemakerUserProfile#sharing_settings}
	SharingSettings *SagemakerUserProfileUserSettingsSharingSettings `field:"optional" json:"sharingSettings" yaml:"sharingSettings"`
	// tensor_board_app_settings block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/sagemaker_user_profile#tensor_board_app_settings SagemakerUserProfile#tensor_board_app_settings}
	TensorBoardAppSettings *SagemakerUserProfileUserSettingsTensorBoardAppSettings `field:"optional" json:"tensorBoardAppSettings" yaml:"tensorBoardAppSettings"`
}

