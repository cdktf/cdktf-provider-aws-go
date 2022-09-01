package kinesis


type Kinesisanalyticsv2ApplicationApplicationConfigurationRunConfiguration struct {
	// application_restore_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/kinesisanalyticsv2_application#application_restore_configuration Kinesisanalyticsv2Application#application_restore_configuration}
	ApplicationRestoreConfiguration *Kinesisanalyticsv2ApplicationApplicationConfigurationRunConfigurationApplicationRestoreConfiguration `field:"optional" json:"applicationRestoreConfiguration" yaml:"applicationRestoreConfiguration"`
	// flink_run_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/kinesisanalyticsv2_application#flink_run_configuration Kinesisanalyticsv2Application#flink_run_configuration}
	FlinkRunConfiguration *Kinesisanalyticsv2ApplicationApplicationConfigurationRunConfigurationFlinkRunConfiguration `field:"optional" json:"flinkRunConfiguration" yaml:"flinkRunConfiguration"`
}

