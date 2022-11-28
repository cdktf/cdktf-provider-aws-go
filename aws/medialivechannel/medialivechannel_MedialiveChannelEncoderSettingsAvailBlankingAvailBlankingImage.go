package medialivechannel


type MedialiveChannelEncoderSettingsAvailBlankingAvailBlankingImage struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/medialive_channel#uri MedialiveChannel#uri}.
	Uri *string `field:"required" json:"uri" yaml:"uri"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/medialive_channel#password_param MedialiveChannel#password_param}.
	PasswordParam *string `field:"optional" json:"passwordParam" yaml:"passwordParam"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/medialive_channel#username MedialiveChannel#username}.
	Username *string `field:"optional" json:"username" yaml:"username"`
}

