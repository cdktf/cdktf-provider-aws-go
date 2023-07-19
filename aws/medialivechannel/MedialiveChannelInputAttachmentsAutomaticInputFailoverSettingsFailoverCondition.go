package medialivechannel


type MedialiveChannelInputAttachmentsAutomaticInputFailoverSettingsFailoverCondition struct {
	// failover_condition_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.8.0/docs/resources/medialive_channel#failover_condition_settings MedialiveChannel#failover_condition_settings}
	FailoverConditionSettings *MedialiveChannelInputAttachmentsAutomaticInputFailoverSettingsFailoverConditionFailoverConditionSettings `field:"optional" json:"failoverConditionSettings" yaml:"failoverConditionSettings"`
}

