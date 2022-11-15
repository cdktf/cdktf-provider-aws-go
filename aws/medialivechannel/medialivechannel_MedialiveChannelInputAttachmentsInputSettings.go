package medialivechannel


type MedialiveChannelInputAttachmentsInputSettings struct {
	// audio_selector block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/medialive_channel#audio_selector MedialiveChannel#audio_selector}
	AudioSelector interface{} `field:"optional" json:"audioSelector" yaml:"audioSelector"`
	// caption_selector block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/medialive_channel#caption_selector MedialiveChannel#caption_selector}
	CaptionSelector interface{} `field:"optional" json:"captionSelector" yaml:"captionSelector"`
}

