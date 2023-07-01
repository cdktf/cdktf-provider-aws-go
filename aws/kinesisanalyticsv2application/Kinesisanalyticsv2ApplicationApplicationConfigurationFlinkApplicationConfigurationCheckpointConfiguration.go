package kinesisanalyticsv2application


type Kinesisanalyticsv2ApplicationApplicationConfigurationFlinkApplicationConfigurationCheckpointConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/resources/kinesisanalyticsv2_application#configuration_type Kinesisanalyticsv2Application#configuration_type}.
	ConfigurationType *string `field:"required" json:"configurationType" yaml:"configurationType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/resources/kinesisanalyticsv2_application#checkpointing_enabled Kinesisanalyticsv2Application#checkpointing_enabled}.
	CheckpointingEnabled interface{} `field:"optional" json:"checkpointingEnabled" yaml:"checkpointingEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/resources/kinesisanalyticsv2_application#checkpoint_interval Kinesisanalyticsv2Application#checkpoint_interval}.
	CheckpointInterval *float64 `field:"optional" json:"checkpointInterval" yaml:"checkpointInterval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/resources/kinesisanalyticsv2_application#min_pause_between_checkpoints Kinesisanalyticsv2Application#min_pause_between_checkpoints}.
	MinPauseBetweenCheckpoints *float64 `field:"optional" json:"minPauseBetweenCheckpoints" yaml:"minPauseBetweenCheckpoints"`
}

