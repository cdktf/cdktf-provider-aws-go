package kinesis


type Kinesisanalyticsv2ApplicationApplicationConfigurationRunConfigurationApplicationRestoreConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/kinesisanalyticsv2_application#application_restore_type Kinesisanalyticsv2Application#application_restore_type}.
	ApplicationRestoreType *string `field:"optional" json:"applicationRestoreType" yaml:"applicationRestoreType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/kinesisanalyticsv2_application#snapshot_name Kinesisanalyticsv2Application#snapshot_name}.
	SnapshotName *string `field:"optional" json:"snapshotName" yaml:"snapshotName"`
}

