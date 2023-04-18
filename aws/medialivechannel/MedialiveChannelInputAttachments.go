package medialivechannel


type MedialiveChannelInputAttachments struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.63.0/docs/resources/medialive_channel#input_attachment_name MedialiveChannel#input_attachment_name}.
	InputAttachmentName *string `field:"required" json:"inputAttachmentName" yaml:"inputAttachmentName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.63.0/docs/resources/medialive_channel#input_id MedialiveChannel#input_id}.
	InputId *string `field:"required" json:"inputId" yaml:"inputId"`
	// automatic_input_failover_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.63.0/docs/resources/medialive_channel#automatic_input_failover_settings MedialiveChannel#automatic_input_failover_settings}
	AutomaticInputFailoverSettings *MedialiveChannelInputAttachmentsAutomaticInputFailoverSettings `field:"optional" json:"automaticInputFailoverSettings" yaml:"automaticInputFailoverSettings"`
	// input_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.63.0/docs/resources/medialive_channel#input_settings MedialiveChannel#input_settings}
	InputSettings *MedialiveChannelInputAttachmentsInputSettings `field:"optional" json:"inputSettings" yaml:"inputSettings"`
}

