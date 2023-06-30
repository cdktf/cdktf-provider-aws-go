package sagemakerdomain


type SagemakerDomainDefaultUserSettingsCanvasAppSettings struct {
	// model_register_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.1/docs/resources/sagemaker_domain#model_register_settings SagemakerDomain#model_register_settings}
	ModelRegisterSettings *SagemakerDomainDefaultUserSettingsCanvasAppSettingsModelRegisterSettings `field:"optional" json:"modelRegisterSettings" yaml:"modelRegisterSettings"`
	// time_series_forecasting_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.1/docs/resources/sagemaker_domain#time_series_forecasting_settings SagemakerDomain#time_series_forecasting_settings}
	TimeSeriesForecastingSettings *SagemakerDomainDefaultUserSettingsCanvasAppSettingsTimeSeriesForecastingSettings `field:"optional" json:"timeSeriesForecastingSettings" yaml:"timeSeriesForecastingSettings"`
}

