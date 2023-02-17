package medialivechannel


type MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorSelectorSettingsTeletextSourceSettings struct {
	// output_rectangle block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/medialive_channel#output_rectangle MedialiveChannel#output_rectangle}
	OutputRectangle *MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorSelectorSettingsTeletextSourceSettingsOutputRectangle `field:"optional" json:"outputRectangle" yaml:"outputRectangle"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/medialive_channel#page_number MedialiveChannel#page_number}.
	PageNumber *string `field:"optional" json:"pageNumber" yaml:"pageNumber"`
}

