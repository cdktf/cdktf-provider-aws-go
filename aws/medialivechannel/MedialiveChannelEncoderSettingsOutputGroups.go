package medialivechannel


type MedialiveChannelEncoderSettingsOutputGroups struct {
	// output_group_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.12.0/docs/resources/medialive_channel#output_group_settings MedialiveChannel#output_group_settings}
	OutputGroupSettings *MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettings `field:"required" json:"outputGroupSettings" yaml:"outputGroupSettings"`
	// outputs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.12.0/docs/resources/medialive_channel#outputs MedialiveChannel#outputs}
	Outputs interface{} `field:"required" json:"outputs" yaml:"outputs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.12.0/docs/resources/medialive_channel#name MedialiveChannel#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

