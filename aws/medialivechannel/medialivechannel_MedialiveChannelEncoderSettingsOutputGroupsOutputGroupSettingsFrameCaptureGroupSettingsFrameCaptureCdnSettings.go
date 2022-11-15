package medialivechannel


type MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsFrameCaptureGroupSettingsFrameCaptureCdnSettings struct {
	// frame_capture_s3_settings block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/medialive_channel#frame_capture_s3_settings MedialiveChannel#frame_capture_s3_settings}
	FrameCaptureS3Settings *MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsFrameCaptureGroupSettingsFrameCaptureCdnSettingsFrameCaptureS3Settings `field:"optional" json:"frameCaptureS3Settings" yaml:"frameCaptureS3Settings"`
}

