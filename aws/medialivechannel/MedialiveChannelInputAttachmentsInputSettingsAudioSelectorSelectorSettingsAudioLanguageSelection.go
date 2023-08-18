package medialivechannel


type MedialiveChannelInputAttachmentsInputSettingsAudioSelectorSelectorSettingsAudioLanguageSelection struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/medialive_channel#language_code MedialiveChannel#language_code}.
	LanguageCode *string `field:"required" json:"languageCode" yaml:"languageCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/medialive_channel#language_selection_policy MedialiveChannel#language_selection_policy}.
	LanguageSelectionPolicy *string `field:"optional" json:"languageSelectionPolicy" yaml:"languageSelectionPolicy"`
}

