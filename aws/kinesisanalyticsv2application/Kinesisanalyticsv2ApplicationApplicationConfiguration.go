package kinesisanalyticsv2application


type Kinesisanalyticsv2ApplicationApplicationConfiguration struct {
	// application_code_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.67.0/docs/resources/kinesisanalyticsv2_application#application_code_configuration Kinesisanalyticsv2Application#application_code_configuration}
	ApplicationCodeConfiguration *Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationCodeConfiguration `field:"required" json:"applicationCodeConfiguration" yaml:"applicationCodeConfiguration"`
	// application_snapshot_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.67.0/docs/resources/kinesisanalyticsv2_application#application_snapshot_configuration Kinesisanalyticsv2Application#application_snapshot_configuration}
	ApplicationSnapshotConfiguration *Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationSnapshotConfiguration `field:"optional" json:"applicationSnapshotConfiguration" yaml:"applicationSnapshotConfiguration"`
	// environment_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.67.0/docs/resources/kinesisanalyticsv2_application#environment_properties Kinesisanalyticsv2Application#environment_properties}
	EnvironmentProperties *Kinesisanalyticsv2ApplicationApplicationConfigurationEnvironmentProperties `field:"optional" json:"environmentProperties" yaml:"environmentProperties"`
	// flink_application_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.67.0/docs/resources/kinesisanalyticsv2_application#flink_application_configuration Kinesisanalyticsv2Application#flink_application_configuration}
	FlinkApplicationConfiguration *Kinesisanalyticsv2ApplicationApplicationConfigurationFlinkApplicationConfiguration `field:"optional" json:"flinkApplicationConfiguration" yaml:"flinkApplicationConfiguration"`
	// run_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.67.0/docs/resources/kinesisanalyticsv2_application#run_configuration Kinesisanalyticsv2Application#run_configuration}
	RunConfiguration *Kinesisanalyticsv2ApplicationApplicationConfigurationRunConfiguration `field:"optional" json:"runConfiguration" yaml:"runConfiguration"`
	// sql_application_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.67.0/docs/resources/kinesisanalyticsv2_application#sql_application_configuration Kinesisanalyticsv2Application#sql_application_configuration}
	SqlApplicationConfiguration *Kinesisanalyticsv2ApplicationApplicationConfigurationSqlApplicationConfiguration `field:"optional" json:"sqlApplicationConfiguration" yaml:"sqlApplicationConfiguration"`
	// vpc_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.67.0/docs/resources/kinesisanalyticsv2_application#vpc_configuration Kinesisanalyticsv2Application#vpc_configuration}
	VpcConfiguration *Kinesisanalyticsv2ApplicationApplicationConfigurationVpcConfiguration `field:"optional" json:"vpcConfiguration" yaml:"vpcConfiguration"`
}

