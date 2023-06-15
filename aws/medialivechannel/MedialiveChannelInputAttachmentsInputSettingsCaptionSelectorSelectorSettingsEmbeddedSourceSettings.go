package medialivechannel


type MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorSelectorSettingsEmbeddedSourceSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.67.0/docs/resources/medialive_channel#convert_608_to_708 MedialiveChannel#convert_608_to_708}.
	Convert608To708 *string `field:"optional" json:"convert608To708" yaml:"convert608To708"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.67.0/docs/resources/medialive_channel#scte20_detection MedialiveChannel#scte20_detection}.
	Scte20Detection *string `field:"optional" json:"scte20Detection" yaml:"scte20Detection"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.67.0/docs/resources/medialive_channel#source_608_channel_number MedialiveChannel#source_608_channel_number}.
	Source608ChannelNumber *float64 `field:"optional" json:"source608ChannelNumber" yaml:"source608ChannelNumber"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.67.0/docs/resources/medialive_channel#source_608_track_number MedialiveChannel#source_608_track_number}.
	Source608TrackNumber *float64 `field:"optional" json:"source608TrackNumber" yaml:"source608TrackNumber"`
}

