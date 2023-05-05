package kinesisanalyticsv2application


type Kinesisanalyticsv2ApplicationApplicationConfigurationFlinkApplicationConfiguration struct {
	// checkpoint_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/kinesisanalyticsv2_application#checkpoint_configuration Kinesisanalyticsv2Application#checkpoint_configuration}
	CheckpointConfiguration *Kinesisanalyticsv2ApplicationApplicationConfigurationFlinkApplicationConfigurationCheckpointConfiguration `field:"optional" json:"checkpointConfiguration" yaml:"checkpointConfiguration"`
	// monitoring_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/kinesisanalyticsv2_application#monitoring_configuration Kinesisanalyticsv2Application#monitoring_configuration}
	MonitoringConfiguration *Kinesisanalyticsv2ApplicationApplicationConfigurationFlinkApplicationConfigurationMonitoringConfiguration `field:"optional" json:"monitoringConfiguration" yaml:"monitoringConfiguration"`
	// parallelism_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/kinesisanalyticsv2_application#parallelism_configuration Kinesisanalyticsv2Application#parallelism_configuration}
	ParallelismConfiguration *Kinesisanalyticsv2ApplicationApplicationConfigurationFlinkApplicationConfigurationParallelismConfiguration `field:"optional" json:"parallelismConfiguration" yaml:"parallelismConfiguration"`
}

