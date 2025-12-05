// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package medialivechannel


type MedialiveChannelInputAttachmentsAutomaticInputFailoverSettingsFailoverConditionFailoverConditionSettings struct {
	// audio_silence_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/medialive_channel#audio_silence_settings MedialiveChannel#audio_silence_settings}
	AudioSilenceSettings *MedialiveChannelInputAttachmentsAutomaticInputFailoverSettingsFailoverConditionFailoverConditionSettingsAudioSilenceSettings `field:"optional" json:"audioSilenceSettings" yaml:"audioSilenceSettings"`
	// input_loss_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/medialive_channel#input_loss_settings MedialiveChannel#input_loss_settings}
	InputLossSettings *MedialiveChannelInputAttachmentsAutomaticInputFailoverSettingsFailoverConditionFailoverConditionSettingsInputLossSettings `field:"optional" json:"inputLossSettings" yaml:"inputLossSettings"`
	// video_black_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/medialive_channel#video_black_settings MedialiveChannel#video_black_settings}
	VideoBlackSettings *MedialiveChannelInputAttachmentsAutomaticInputFailoverSettingsFailoverConditionFailoverConditionSettingsVideoBlackSettings `field:"optional" json:"videoBlackSettings" yaml:"videoBlackSettings"`
}

