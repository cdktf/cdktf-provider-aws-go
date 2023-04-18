package medialivechannel


type MedialiveChannelEncoderSettingsAudioDescriptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.63.0/docs/resources/medialive_channel#audio_selector_name MedialiveChannel#audio_selector_name}.
	AudioSelectorName *string `field:"required" json:"audioSelectorName" yaml:"audioSelectorName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.63.0/docs/resources/medialive_channel#name MedialiveChannel#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// audio_normalization_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.63.0/docs/resources/medialive_channel#audio_normalization_settings MedialiveChannel#audio_normalization_settings}
	AudioNormalizationSettings *MedialiveChannelEncoderSettingsAudioDescriptionsAudioNormalizationSettings `field:"optional" json:"audioNormalizationSettings" yaml:"audioNormalizationSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.63.0/docs/resources/medialive_channel#audio_type MedialiveChannel#audio_type}.
	AudioType *string `field:"optional" json:"audioType" yaml:"audioType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.63.0/docs/resources/medialive_channel#audio_type_control MedialiveChannel#audio_type_control}.
	AudioTypeControl *string `field:"optional" json:"audioTypeControl" yaml:"audioTypeControl"`
	// audio_watermark_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.63.0/docs/resources/medialive_channel#audio_watermark_settings MedialiveChannel#audio_watermark_settings}
	AudioWatermarkSettings *MedialiveChannelEncoderSettingsAudioDescriptionsAudioWatermarkSettings `field:"optional" json:"audioWatermarkSettings" yaml:"audioWatermarkSettings"`
	// codec_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.63.0/docs/resources/medialive_channel#codec_settings MedialiveChannel#codec_settings}
	CodecSettings *MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettings `field:"optional" json:"codecSettings" yaml:"codecSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.63.0/docs/resources/medialive_channel#language_code MedialiveChannel#language_code}.
	LanguageCode *string `field:"optional" json:"languageCode" yaml:"languageCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.63.0/docs/resources/medialive_channel#language_code_control MedialiveChannel#language_code_control}.
	LanguageCodeControl *string `field:"optional" json:"languageCodeControl" yaml:"languageCodeControl"`
	// remix_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.63.0/docs/resources/medialive_channel#remix_settings MedialiveChannel#remix_settings}
	RemixSettings *MedialiveChannelEncoderSettingsAudioDescriptionsRemixSettings `field:"optional" json:"remixSettings" yaml:"remixSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.63.0/docs/resources/medialive_channel#stream_name MedialiveChannel#stream_name}.
	StreamName *string `field:"optional" json:"streamName" yaml:"streamName"`
}

