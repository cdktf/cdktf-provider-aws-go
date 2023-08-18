package medialivechannel


type MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsBurnInDestinationSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/medialive_channel#outline_color MedialiveChannel#outline_color}.
	OutlineColor *string `field:"required" json:"outlineColor" yaml:"outlineColor"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/medialive_channel#teletext_grid_control MedialiveChannel#teletext_grid_control}.
	TeletextGridControl *string `field:"required" json:"teletextGridControl" yaml:"teletextGridControl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/medialive_channel#alignment MedialiveChannel#alignment}.
	Alignment *string `field:"optional" json:"alignment" yaml:"alignment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/medialive_channel#background_color MedialiveChannel#background_color}.
	BackgroundColor *string `field:"optional" json:"backgroundColor" yaml:"backgroundColor"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/medialive_channel#background_opacity MedialiveChannel#background_opacity}.
	BackgroundOpacity *float64 `field:"optional" json:"backgroundOpacity" yaml:"backgroundOpacity"`
	// font block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/medialive_channel#font MedialiveChannel#font}
	Font *MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsBurnInDestinationSettingsFont `field:"optional" json:"font" yaml:"font"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/medialive_channel#font_color MedialiveChannel#font_color}.
	FontColor *string `field:"optional" json:"fontColor" yaml:"fontColor"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/medialive_channel#font_opacity MedialiveChannel#font_opacity}.
	FontOpacity *float64 `field:"optional" json:"fontOpacity" yaml:"fontOpacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/medialive_channel#font_resolution MedialiveChannel#font_resolution}.
	FontResolution *float64 `field:"optional" json:"fontResolution" yaml:"fontResolution"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/medialive_channel#font_size MedialiveChannel#font_size}.
	FontSize *string `field:"optional" json:"fontSize" yaml:"fontSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/medialive_channel#outline_size MedialiveChannel#outline_size}.
	OutlineSize *float64 `field:"optional" json:"outlineSize" yaml:"outlineSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/medialive_channel#shadow_color MedialiveChannel#shadow_color}.
	ShadowColor *string `field:"optional" json:"shadowColor" yaml:"shadowColor"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/medialive_channel#shadow_opacity MedialiveChannel#shadow_opacity}.
	ShadowOpacity *float64 `field:"optional" json:"shadowOpacity" yaml:"shadowOpacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/medialive_channel#shadow_x_offset MedialiveChannel#shadow_x_offset}.
	ShadowXOffset *float64 `field:"optional" json:"shadowXOffset" yaml:"shadowXOffset"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/medialive_channel#shadow_y_offset MedialiveChannel#shadow_y_offset}.
	ShadowYOffset *float64 `field:"optional" json:"shadowYOffset" yaml:"shadowYOffset"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/medialive_channel#x_position MedialiveChannel#x_position}.
	XPosition *float64 `field:"optional" json:"xPosition" yaml:"xPosition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/medialive_channel#y_position MedialiveChannel#y_position}.
	YPosition *float64 `field:"optional" json:"yPosition" yaml:"yPosition"`
}

