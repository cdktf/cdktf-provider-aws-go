package medialivechannel


type MedialiveChannelInputAttachmentsInputSettingsNetworkInputSettings struct {
	// hls_input_settings block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/medialive_channel#hls_input_settings MedialiveChannel#hls_input_settings}
	HlsInputSettings *MedialiveChannelInputAttachmentsInputSettingsNetworkInputSettingsHlsInputSettings `field:"optional" json:"hlsInputSettings" yaml:"hlsInputSettings"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/medialive_channel#server_validation MedialiveChannel#server_validation}.
	ServerValidation *string `field:"optional" json:"serverValidation" yaml:"serverValidation"`
}

