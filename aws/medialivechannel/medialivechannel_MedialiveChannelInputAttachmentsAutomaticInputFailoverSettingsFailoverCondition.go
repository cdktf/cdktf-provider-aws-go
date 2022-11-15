package medialivechannel


type MedialiveChannelInputAttachmentsAutomaticInputFailoverSettingsFailoverCondition struct {
	// failover_condition_settings block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/medialive_channel#failover_condition_settings MedialiveChannel#failover_condition_settings}
	FailoverConditionSettings *MedialiveChannelInputAttachmentsAutomaticInputFailoverSettingsFailoverConditionFailoverConditionSettings `field:"optional" json:"failoverConditionSettings" yaml:"failoverConditionSettings"`
}

