package ivsrecordingconfiguration


type IvsRecordingConfigurationThumbnailConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ivs_recording_configuration#recording_mode IvsRecordingConfiguration#recording_mode}.
	RecordingMode *string `field:"optional" json:"recordingMode" yaml:"recordingMode"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ivs_recording_configuration#target_interval_seconds IvsRecordingConfiguration#target_interval_seconds}.
	TargetIntervalSeconds *float64 `field:"optional" json:"targetIntervalSeconds" yaml:"targetIntervalSeconds"`
}

