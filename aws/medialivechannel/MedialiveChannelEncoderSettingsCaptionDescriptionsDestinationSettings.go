// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package medialivechannel


type MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettings struct {
	// arib_destination_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.51.1/docs/resources/medialive_channel#arib_destination_settings MedialiveChannel#arib_destination_settings}
	AribDestinationSettings *MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsAribDestinationSettings `field:"optional" json:"aribDestinationSettings" yaml:"aribDestinationSettings"`
	// burn_in_destination_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.51.1/docs/resources/medialive_channel#burn_in_destination_settings MedialiveChannel#burn_in_destination_settings}
	BurnInDestinationSettings *MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsBurnInDestinationSettings `field:"optional" json:"burnInDestinationSettings" yaml:"burnInDestinationSettings"`
	// dvb_sub_destination_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.51.1/docs/resources/medialive_channel#dvb_sub_destination_settings MedialiveChannel#dvb_sub_destination_settings}
	DvbSubDestinationSettings *MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettings `field:"optional" json:"dvbSubDestinationSettings" yaml:"dvbSubDestinationSettings"`
	// ebu_tt_d_destination_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.51.1/docs/resources/medialive_channel#ebu_tt_d_destination_settings MedialiveChannel#ebu_tt_d_destination_settings}
	EbuTtDDestinationSettings *MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsEbuTtDDestinationSettings `field:"optional" json:"ebuTtDDestinationSettings" yaml:"ebuTtDDestinationSettings"`
	// embedded_destination_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.51.1/docs/resources/medialive_channel#embedded_destination_settings MedialiveChannel#embedded_destination_settings}
	EmbeddedDestinationSettings *MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsEmbeddedDestinationSettings `field:"optional" json:"embeddedDestinationSettings" yaml:"embeddedDestinationSettings"`
	// embedded_plus_scte20_destination_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.51.1/docs/resources/medialive_channel#embedded_plus_scte20_destination_settings MedialiveChannel#embedded_plus_scte20_destination_settings}
	EmbeddedPlusScte20DestinationSettings *MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsEmbeddedPlusScte20DestinationSettings `field:"optional" json:"embeddedPlusScte20DestinationSettings" yaml:"embeddedPlusScte20DestinationSettings"`
	// rtmp_caption_info_destination_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.51.1/docs/resources/medialive_channel#rtmp_caption_info_destination_settings MedialiveChannel#rtmp_caption_info_destination_settings}
	RtmpCaptionInfoDestinationSettings *MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsRtmpCaptionInfoDestinationSettings `field:"optional" json:"rtmpCaptionInfoDestinationSettings" yaml:"rtmpCaptionInfoDestinationSettings"`
	// scte20_plus_embedded_destination_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.51.1/docs/resources/medialive_channel#scte20_plus_embedded_destination_settings MedialiveChannel#scte20_plus_embedded_destination_settings}
	Scte20PlusEmbeddedDestinationSettings *MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsScte20PlusEmbeddedDestinationSettings `field:"optional" json:"scte20PlusEmbeddedDestinationSettings" yaml:"scte20PlusEmbeddedDestinationSettings"`
	// scte27_destination_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.51.1/docs/resources/medialive_channel#scte27_destination_settings MedialiveChannel#scte27_destination_settings}
	Scte27DestinationSettings *MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsScte27DestinationSettings `field:"optional" json:"scte27DestinationSettings" yaml:"scte27DestinationSettings"`
	// smpte_tt_destination_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.51.1/docs/resources/medialive_channel#smpte_tt_destination_settings MedialiveChannel#smpte_tt_destination_settings}
	SmpteTtDestinationSettings *MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsSmpteTtDestinationSettings `field:"optional" json:"smpteTtDestinationSettings" yaml:"smpteTtDestinationSettings"`
	// teletext_destination_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.51.1/docs/resources/medialive_channel#teletext_destination_settings MedialiveChannel#teletext_destination_settings}
	TeletextDestinationSettings *MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsTeletextDestinationSettings `field:"optional" json:"teletextDestinationSettings" yaml:"teletextDestinationSettings"`
	// ttml_destination_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.51.1/docs/resources/medialive_channel#ttml_destination_settings MedialiveChannel#ttml_destination_settings}
	TtmlDestinationSettings *MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsTtmlDestinationSettings `field:"optional" json:"ttmlDestinationSettings" yaml:"ttmlDestinationSettings"`
	// webvtt_destination_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.51.1/docs/resources/medialive_channel#webvtt_destination_settings MedialiveChannel#webvtt_destination_settings}
	WebvttDestinationSettings *MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsWebvttDestinationSettings `field:"optional" json:"webvttDestinationSettings" yaml:"webvttDestinationSettings"`
}

