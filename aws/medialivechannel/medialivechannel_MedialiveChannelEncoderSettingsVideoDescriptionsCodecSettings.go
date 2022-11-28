package medialivechannel


type MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettings struct {
	// frame_capture_settings block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/medialive_channel#frame_capture_settings MedialiveChannel#frame_capture_settings}
	FrameCaptureSettings *MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsFrameCaptureSettings `field:"optional" json:"frameCaptureSettings" yaml:"frameCaptureSettings"`
	// h264_settings block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/medialive_channel#h264_settings MedialiveChannel#h264_settings}
	H264Settings *MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264Settings `field:"optional" json:"h264Settings" yaml:"h264Settings"`
}

