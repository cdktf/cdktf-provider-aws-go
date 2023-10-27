// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakerdomain


type SagemakerDomainDefaultUserSettingsCanvasAppSettings struct {
	// model_register_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/sagemaker_domain#model_register_settings SagemakerDomain#model_register_settings}
	ModelRegisterSettings *SagemakerDomainDefaultUserSettingsCanvasAppSettingsModelRegisterSettings `field:"optional" json:"modelRegisterSettings" yaml:"modelRegisterSettings"`
	// time_series_forecasting_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/sagemaker_domain#time_series_forecasting_settings SagemakerDomain#time_series_forecasting_settings}
	TimeSeriesForecastingSettings *SagemakerDomainDefaultUserSettingsCanvasAppSettingsTimeSeriesForecastingSettings `field:"optional" json:"timeSeriesForecastingSettings" yaml:"timeSeriesForecastingSettings"`
	// workspace_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/sagemaker_domain#workspace_settings SagemakerDomain#workspace_settings}
	WorkspaceSettings *SagemakerDomainDefaultUserSettingsCanvasAppSettingsWorkspaceSettings `field:"optional" json:"workspaceSettings" yaml:"workspaceSettings"`
}

