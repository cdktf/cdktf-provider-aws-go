package medialivechannel


type MedialiveChannelEncoderSettingsAvailBlanking struct {
	// avail_blanking_image block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/medialive_channel#avail_blanking_image MedialiveChannel#avail_blanking_image}
	AvailBlankingImage *MedialiveChannelEncoderSettingsAvailBlankingAvailBlankingImage `field:"optional" json:"availBlankingImage" yaml:"availBlankingImage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/medialive_channel#state MedialiveChannel#state}.
	State *string `field:"optional" json:"state" yaml:"state"`
}

