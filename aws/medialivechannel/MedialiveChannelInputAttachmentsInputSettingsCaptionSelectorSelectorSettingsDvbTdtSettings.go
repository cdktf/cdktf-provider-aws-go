package medialivechannel


type MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorSelectorSettingsDvbTdtSettings struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/medialive_channel#ocr_language MedialiveChannel#ocr_language}.
	OcrLanguage *string `field:"optional" json:"ocrLanguage" yaml:"ocrLanguage"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/medialive_channel#pid MedialiveChannel#pid}.
	Pid *float64 `field:"optional" json:"pid" yaml:"pid"`
}

