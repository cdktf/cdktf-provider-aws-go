package medialivechannel


type MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsFrameCaptureSettings struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/medialive_channel#capture_interval MedialiveChannel#capture_interval}.
	CaptureInterval *float64 `field:"optional" json:"captureInterval" yaml:"captureInterval"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/medialive_channel#capture_interval_units MedialiveChannel#capture_interval_units}.
	CaptureIntervalUnits *string `field:"optional" json:"captureIntervalUnits" yaml:"captureIntervalUnits"`
}

