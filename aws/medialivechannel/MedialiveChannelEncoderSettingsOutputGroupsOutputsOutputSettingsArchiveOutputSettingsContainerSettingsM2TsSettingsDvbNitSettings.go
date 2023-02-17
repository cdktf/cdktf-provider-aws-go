package medialivechannel


type MedialiveChannelEncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2TsSettingsDvbNitSettings struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/medialive_channel#network_id MedialiveChannel#network_id}.
	NetworkId *float64 `field:"required" json:"networkId" yaml:"networkId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/medialive_channel#network_name MedialiveChannel#network_name}.
	NetworkName *string `field:"required" json:"networkName" yaml:"networkName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/medialive_channel#rep_interval MedialiveChannel#rep_interval}.
	RepInterval *float64 `field:"optional" json:"repInterval" yaml:"repInterval"`
}

