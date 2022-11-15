package medialivechannel


type MedialiveChannelEncoderSettingsOutputGroupsOutputsOutputSettingsMsSmoothOutputSettings struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/medialive_channel#h265_packaging_type MedialiveChannel#h265_packaging_type}.
	H265PackagingType *string `field:"optional" json:"h265PackagingType" yaml:"h265PackagingType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/medialive_channel#name_modifier MedialiveChannel#name_modifier}.
	NameModifier *string `field:"optional" json:"nameModifier" yaml:"nameModifier"`
}

