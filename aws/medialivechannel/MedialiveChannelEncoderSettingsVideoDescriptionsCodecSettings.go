package medialivechannel


type MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettings struct {
	// frame_capture_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/resources/medialive_channel#frame_capture_settings MedialiveChannel#frame_capture_settings}
	FrameCaptureSettings *MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsFrameCaptureSettings `field:"optional" json:"frameCaptureSettings" yaml:"frameCaptureSettings"`
	// h264_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/resources/medialive_channel#h264_settings MedialiveChannel#h264_settings}
	H264Settings *MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264Settings `field:"optional" json:"h264Settings" yaml:"h264Settings"`
}

