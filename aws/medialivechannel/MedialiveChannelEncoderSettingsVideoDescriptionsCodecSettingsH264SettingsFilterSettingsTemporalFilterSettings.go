package medialivechannel


type MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsFilterSettingsTemporalFilterSettings struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/medialive_channel#post_filter_sharpening MedialiveChannel#post_filter_sharpening}.
	PostFilterSharpening *string `field:"optional" json:"postFilterSharpening" yaml:"postFilterSharpening"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/medialive_channel#strength MedialiveChannel#strength}.
	Strength *string `field:"optional" json:"strength" yaml:"strength"`
}

