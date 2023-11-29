// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakeruserprofile


type SagemakerUserProfileUserSettingsCanvasAppSettings struct {
	// direct_deploy_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/sagemaker_user_profile#direct_deploy_settings SagemakerUserProfile#direct_deploy_settings}
	DirectDeploySettings *SagemakerUserProfileUserSettingsCanvasAppSettingsDirectDeploySettings `field:"optional" json:"directDeploySettings" yaml:"directDeploySettings"`
	// identity_provider_oauth_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/sagemaker_user_profile#identity_provider_oauth_settings SagemakerUserProfile#identity_provider_oauth_settings}
	IdentityProviderOauthSettings interface{} `field:"optional" json:"identityProviderOauthSettings" yaml:"identityProviderOauthSettings"`
	// kendra_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/sagemaker_user_profile#kendra_settings SagemakerUserProfile#kendra_settings}
	KendraSettings *SagemakerUserProfileUserSettingsCanvasAppSettingsKendraSettings `field:"optional" json:"kendraSettings" yaml:"kendraSettings"`
	// model_register_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/sagemaker_user_profile#model_register_settings SagemakerUserProfile#model_register_settings}
	ModelRegisterSettings *SagemakerUserProfileUserSettingsCanvasAppSettingsModelRegisterSettings `field:"optional" json:"modelRegisterSettings" yaml:"modelRegisterSettings"`
	// time_series_forecasting_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/sagemaker_user_profile#time_series_forecasting_settings SagemakerUserProfile#time_series_forecasting_settings}
	TimeSeriesForecastingSettings *SagemakerUserProfileUserSettingsCanvasAppSettingsTimeSeriesForecastingSettings `field:"optional" json:"timeSeriesForecastingSettings" yaml:"timeSeriesForecastingSettings"`
	// workspace_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/sagemaker_user_profile#workspace_settings SagemakerUserProfile#workspace_settings}
	WorkspaceSettings *SagemakerUserProfileUserSettingsCanvasAppSettingsWorkspaceSettings `field:"optional" json:"workspaceSettings" yaml:"workspaceSettings"`
}

