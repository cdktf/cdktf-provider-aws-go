package medialivechannel


type MedialiveChannelEncoderSettingsAudioDescriptionsRemixSettingsChannelMappings struct {
	// input_channel_levels block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.63.0/docs/resources/medialive_channel#input_channel_levels MedialiveChannel#input_channel_levels}
	InputChannelLevels interface{} `field:"required" json:"inputChannelLevels" yaml:"inputChannelLevels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.63.0/docs/resources/medialive_channel#output_channel MedialiveChannel#output_channel}.
	OutputChannel *float64 `field:"required" json:"outputChannel" yaml:"outputChannel"`
}

