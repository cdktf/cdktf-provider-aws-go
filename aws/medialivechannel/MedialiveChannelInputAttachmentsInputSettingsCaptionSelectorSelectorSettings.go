package medialivechannel


type MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorSelectorSettings struct {
	// ancillary_source_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/medialive_channel#ancillary_source_settings MedialiveChannel#ancillary_source_settings}
	AncillarySourceSettings *MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorSelectorSettingsAncillarySourceSettings `field:"optional" json:"ancillarySourceSettings" yaml:"ancillarySourceSettings"`
	// dvb_tdt_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/medialive_channel#dvb_tdt_settings MedialiveChannel#dvb_tdt_settings}
	DvbTdtSettings *MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorSelectorSettingsDvbTdtSettings `field:"optional" json:"dvbTdtSettings" yaml:"dvbTdtSettings"`
	// embedded_source_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/medialive_channel#embedded_source_settings MedialiveChannel#embedded_source_settings}
	EmbeddedSourceSettings *MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorSelectorSettingsEmbeddedSourceSettings `field:"optional" json:"embeddedSourceSettings" yaml:"embeddedSourceSettings"`
	// scte20_source_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/medialive_channel#scte20_source_settings MedialiveChannel#scte20_source_settings}
	Scte20SourceSettings *MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorSelectorSettingsScte20SourceSettings `field:"optional" json:"scte20SourceSettings" yaml:"scte20SourceSettings"`
	// scte27_source_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/medialive_channel#scte27_source_settings MedialiveChannel#scte27_source_settings}
	Scte27SourceSettings *MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorSelectorSettingsScte27SourceSettings `field:"optional" json:"scte27SourceSettings" yaml:"scte27SourceSettings"`
	// teletext_source_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/medialive_channel#teletext_source_settings MedialiveChannel#teletext_source_settings}
	TeletextSourceSettings *MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorSelectorSettingsTeletextSourceSettings `field:"optional" json:"teletextSourceSettings" yaml:"teletextSourceSettings"`
}

