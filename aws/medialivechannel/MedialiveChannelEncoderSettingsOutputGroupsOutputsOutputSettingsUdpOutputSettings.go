package medialivechannel


type MedialiveChannelEncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettings struct {
	// container_settings block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/medialive_channel#container_settings MedialiveChannel#container_settings}
	ContainerSettings *MedialiveChannelEncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsContainerSettings `field:"required" json:"containerSettings" yaml:"containerSettings"`
	// destination block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/medialive_channel#destination MedialiveChannel#destination}
	Destination *MedialiveChannelEncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsDestination `field:"required" json:"destination" yaml:"destination"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/medialive_channel#buffer_msec MedialiveChannel#buffer_msec}.
	BufferMsec *float64 `field:"optional" json:"bufferMsec" yaml:"bufferMsec"`
	// fec_output_settings block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/medialive_channel#fec_output_settings MedialiveChannel#fec_output_settings}
	FecOutputSettings *MedialiveChannelEncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsFecOutputSettings `field:"optional" json:"fecOutputSettings" yaml:"fecOutputSettings"`
}

