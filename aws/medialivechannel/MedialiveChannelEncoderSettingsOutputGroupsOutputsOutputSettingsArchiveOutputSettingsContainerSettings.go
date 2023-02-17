package medialivechannel


type MedialiveChannelEncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettings struct {
	// m2ts_settings block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/medialive_channel#m2ts_settings MedialiveChannel#m2ts_settings}
	M2TsSettings *MedialiveChannelEncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2TsSettings `field:"optional" json:"m2TsSettings" yaml:"m2TsSettings"`
	// raw_settings block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/medialive_channel#raw_settings MedialiveChannel#raw_settings}
	RawSettings *MedialiveChannelEncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsRawSettings `field:"optional" json:"rawSettings" yaml:"rawSettings"`
}

