package medialivechannel


type MedialiveChannelEncoderSettingsMotionGraphicsConfiguration struct {
	// motion_graphics_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/medialive_channel#motion_graphics_settings MedialiveChannel#motion_graphics_settings}
	MotionGraphicsSettings *MedialiveChannelEncoderSettingsMotionGraphicsConfigurationMotionGraphicsSettings `field:"required" json:"motionGraphicsSettings" yaml:"motionGraphicsSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/medialive_channel#motion_graphics_insertion MedialiveChannel#motion_graphics_insertion}.
	MotionGraphicsInsertion *string `field:"optional" json:"motionGraphicsInsertion" yaml:"motionGraphicsInsertion"`
}

