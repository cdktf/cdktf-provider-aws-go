package medialivechannel


type MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsFrameCaptureGroupSettings struct {
	// destination block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/medialive_channel#destination MedialiveChannel#destination}
	Destination *MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsFrameCaptureGroupSettingsDestination `field:"required" json:"destination" yaml:"destination"`
	// frame_capture_cdn_settings block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/medialive_channel#frame_capture_cdn_settings MedialiveChannel#frame_capture_cdn_settings}
	FrameCaptureCdnSettings *MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsFrameCaptureGroupSettingsFrameCaptureCdnSettings `field:"optional" json:"frameCaptureCdnSettings" yaml:"frameCaptureCdnSettings"`
}

