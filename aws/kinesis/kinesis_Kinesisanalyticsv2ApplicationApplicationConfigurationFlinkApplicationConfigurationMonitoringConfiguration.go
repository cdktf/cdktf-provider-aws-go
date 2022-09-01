package kinesis


type Kinesisanalyticsv2ApplicationApplicationConfigurationFlinkApplicationConfigurationMonitoringConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/kinesisanalyticsv2_application#configuration_type Kinesisanalyticsv2Application#configuration_type}.
	ConfigurationType *string `field:"required" json:"configurationType" yaml:"configurationType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/kinesisanalyticsv2_application#log_level Kinesisanalyticsv2Application#log_level}.
	LogLevel *string `field:"optional" json:"logLevel" yaml:"logLevel"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/kinesisanalyticsv2_application#metrics_level Kinesisanalyticsv2Application#metrics_level}.
	MetricsLevel *string `field:"optional" json:"metricsLevel" yaml:"metricsLevel"`
}

