package medialivechannel


type MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsRtmpGroupSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.67.0/docs/resources/medialive_channel#ad_markers MedialiveChannel#ad_markers}.
	AdMarkers *[]*string `field:"optional" json:"adMarkers" yaml:"adMarkers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.67.0/docs/resources/medialive_channel#authentication_scheme MedialiveChannel#authentication_scheme}.
	AuthenticationScheme *string `field:"optional" json:"authenticationScheme" yaml:"authenticationScheme"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.67.0/docs/resources/medialive_channel#cache_full_behavior MedialiveChannel#cache_full_behavior}.
	CacheFullBehavior *string `field:"optional" json:"cacheFullBehavior" yaml:"cacheFullBehavior"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.67.0/docs/resources/medialive_channel#cache_length MedialiveChannel#cache_length}.
	CacheLength *float64 `field:"optional" json:"cacheLength" yaml:"cacheLength"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.67.0/docs/resources/medialive_channel#caption_data MedialiveChannel#caption_data}.
	CaptionData *string `field:"optional" json:"captionData" yaml:"captionData"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.67.0/docs/resources/medialive_channel#input_loss_action MedialiveChannel#input_loss_action}.
	InputLossAction *string `field:"optional" json:"inputLossAction" yaml:"inputLossAction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.67.0/docs/resources/medialive_channel#restart_delay MedialiveChannel#restart_delay}.
	RestartDelay *float64 `field:"optional" json:"restartDelay" yaml:"restartDelay"`
}

