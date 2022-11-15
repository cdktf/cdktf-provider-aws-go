package medialivechannel


type MedialiveChannelInputAttachmentsInputSettingsCaptionSelector struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/medialive_channel#name MedialiveChannel#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/medialive_channel#deblock_filter MedialiveChannel#deblock_filter}.
	DeblockFilter *string `field:"optional" json:"deblockFilter" yaml:"deblockFilter"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/medialive_channel#denoise_filter MedialiveChannel#denoise_filter}.
	DenoiseFilter *string `field:"optional" json:"denoiseFilter" yaml:"denoiseFilter"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/medialive_channel#filter_strength MedialiveChannel#filter_strength}.
	FilterStrength *float64 `field:"optional" json:"filterStrength" yaml:"filterStrength"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/medialive_channel#input_filter MedialiveChannel#input_filter}.
	InputFilter *string `field:"optional" json:"inputFilter" yaml:"inputFilter"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/medialive_channel#language_code MedialiveChannel#language_code}.
	LanguageCode *string `field:"optional" json:"languageCode" yaml:"languageCode"`
	// network_input_settings block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/medialive_channel#network_input_settings MedialiveChannel#network_input_settings}
	NetworkInputSettings *MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorNetworkInputSettings `field:"optional" json:"networkInputSettings" yaml:"networkInputSettings"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/medialive_channel#scte_35_pid MedialiveChannel#scte_35_pid}.
	Scte35Pid *float64 `field:"optional" json:"scte35Pid" yaml:"scte35Pid"`
	// selector_settings block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/medialive_channel#selector_settings MedialiveChannel#selector_settings}
	SelectorSettings *MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorSelectorSettings `field:"optional" json:"selectorSettings" yaml:"selectorSettings"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/medialive_channel#smpte_2038_data_preference MedialiveChannel#smpte_2038_data_preference}.
	Smpte2038DataPreference *string `field:"optional" json:"smpte2038DataPreference" yaml:"smpte2038DataPreference"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/medialive_channel#source_end_behavior MedialiveChannel#source_end_behavior}.
	SourceEndBehavior *string `field:"optional" json:"sourceEndBehavior" yaml:"sourceEndBehavior"`
	// video_selector block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/medialive_channel#video_selector MedialiveChannel#video_selector}
	VideoSelector *MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorVideoSelector `field:"optional" json:"videoSelector" yaml:"videoSelector"`
}

