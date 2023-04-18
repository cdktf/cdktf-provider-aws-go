package medialivechannel


type MedialiveChannelInputAttachmentsInputSettings struct {
	// audio_selector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.63.0/docs/resources/medialive_channel#audio_selector MedialiveChannel#audio_selector}
	AudioSelector interface{} `field:"optional" json:"audioSelector" yaml:"audioSelector"`
	// caption_selector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.63.0/docs/resources/medialive_channel#caption_selector MedialiveChannel#caption_selector}
	CaptionSelector interface{} `field:"optional" json:"captionSelector" yaml:"captionSelector"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.63.0/docs/resources/medialive_channel#deblock_filter MedialiveChannel#deblock_filter}.
	DeblockFilter *string `field:"optional" json:"deblockFilter" yaml:"deblockFilter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.63.0/docs/resources/medialive_channel#denoise_filter MedialiveChannel#denoise_filter}.
	DenoiseFilter *string `field:"optional" json:"denoiseFilter" yaml:"denoiseFilter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.63.0/docs/resources/medialive_channel#filter_strength MedialiveChannel#filter_strength}.
	FilterStrength *float64 `field:"optional" json:"filterStrength" yaml:"filterStrength"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.63.0/docs/resources/medialive_channel#input_filter MedialiveChannel#input_filter}.
	InputFilter *string `field:"optional" json:"inputFilter" yaml:"inputFilter"`
	// network_input_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.63.0/docs/resources/medialive_channel#network_input_settings MedialiveChannel#network_input_settings}
	NetworkInputSettings *MedialiveChannelInputAttachmentsInputSettingsNetworkInputSettings `field:"optional" json:"networkInputSettings" yaml:"networkInputSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.63.0/docs/resources/medialive_channel#scte35_pid MedialiveChannel#scte35_pid}.
	Scte35Pid *float64 `field:"optional" json:"scte35Pid" yaml:"scte35Pid"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.63.0/docs/resources/medialive_channel#smpte2038_data_preference MedialiveChannel#smpte2038_data_preference}.
	Smpte2038DataPreference *string `field:"optional" json:"smpte2038DataPreference" yaml:"smpte2038DataPreference"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.63.0/docs/resources/medialive_channel#source_end_behavior MedialiveChannel#source_end_behavior}.
	SourceEndBehavior *string `field:"optional" json:"sourceEndBehavior" yaml:"sourceEndBehavior"`
	// video_selector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.63.0/docs/resources/medialive_channel#video_selector MedialiveChannel#video_selector}
	VideoSelector *MedialiveChannelInputAttachmentsInputSettingsVideoSelector `field:"optional" json:"videoSelector" yaml:"videoSelector"`
}

