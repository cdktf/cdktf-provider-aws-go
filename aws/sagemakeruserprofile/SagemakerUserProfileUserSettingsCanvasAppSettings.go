package sagemakeruserprofile


type SagemakerUserProfileUserSettingsCanvasAppSettings struct {
	// time_series_forecasting_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.64.0/docs/resources/sagemaker_user_profile#time_series_forecasting_settings SagemakerUserProfile#time_series_forecasting_settings}
	TimeSeriesForecastingSettings *SagemakerUserProfileUserSettingsCanvasAppSettingsTimeSeriesForecastingSettings `field:"optional" json:"timeSeriesForecastingSettings" yaml:"timeSeriesForecastingSettings"`
}

