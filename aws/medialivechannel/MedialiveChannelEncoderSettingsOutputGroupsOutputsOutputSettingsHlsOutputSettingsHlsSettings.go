package medialivechannel


type MedialiveChannelEncoderSettingsOutputGroupsOutputsOutputSettingsHlsOutputSettingsHlsSettings struct {
	// audio_only_hls_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/medialive_channel#audio_only_hls_settings MedialiveChannel#audio_only_hls_settings}
	AudioOnlyHlsSettings *MedialiveChannelEncoderSettingsOutputGroupsOutputsOutputSettingsHlsOutputSettingsHlsSettingsAudioOnlyHlsSettings `field:"optional" json:"audioOnlyHlsSettings" yaml:"audioOnlyHlsSettings"`
	// fmp4_hls_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/medialive_channel#fmp4_hls_settings MedialiveChannel#fmp4_hls_settings}
	Fmp4HlsSettings *MedialiveChannelEncoderSettingsOutputGroupsOutputsOutputSettingsHlsOutputSettingsHlsSettingsFmp4HlsSettings `field:"optional" json:"fmp4HlsSettings" yaml:"fmp4HlsSettings"`
	// frame_capture_hls_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/medialive_channel#frame_capture_hls_settings MedialiveChannel#frame_capture_hls_settings}
	FrameCaptureHlsSettings *MedialiveChannelEncoderSettingsOutputGroupsOutputsOutputSettingsHlsOutputSettingsHlsSettingsFrameCaptureHlsSettings `field:"optional" json:"frameCaptureHlsSettings" yaml:"frameCaptureHlsSettings"`
	// standard_hls_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/medialive_channel#standard_hls_settings MedialiveChannel#standard_hls_settings}
	StandardHlsSettings *MedialiveChannelEncoderSettingsOutputGroupsOutputsOutputSettingsHlsOutputSettingsHlsSettingsStandardHlsSettings `field:"optional" json:"standardHlsSettings" yaml:"standardHlsSettings"`
}

