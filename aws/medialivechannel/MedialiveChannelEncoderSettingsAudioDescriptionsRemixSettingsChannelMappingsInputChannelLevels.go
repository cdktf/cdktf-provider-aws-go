package medialivechannel


type MedialiveChannelEncoderSettingsAudioDescriptionsRemixSettingsChannelMappingsInputChannelLevels struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.10.0/docs/resources/medialive_channel#gain MedialiveChannel#gain}.
	Gain *float64 `field:"required" json:"gain" yaml:"gain"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.10.0/docs/resources/medialive_channel#input_channel MedialiveChannel#input_channel}.
	InputChannel *float64 `field:"required" json:"inputChannel" yaml:"inputChannel"`
}

