// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package medialivechannel


type MedialiveChannelEncoderSettingsGlobalConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.59.0/docs/resources/medialive_channel#initial_audio_gain MedialiveChannel#initial_audio_gain}.
	InitialAudioGain *float64 `field:"optional" json:"initialAudioGain" yaml:"initialAudioGain"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.59.0/docs/resources/medialive_channel#input_end_action MedialiveChannel#input_end_action}.
	InputEndAction *string `field:"optional" json:"inputEndAction" yaml:"inputEndAction"`
	// input_loss_behavior block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.59.0/docs/resources/medialive_channel#input_loss_behavior MedialiveChannel#input_loss_behavior}
	InputLossBehavior *MedialiveChannelEncoderSettingsGlobalConfigurationInputLossBehavior `field:"optional" json:"inputLossBehavior" yaml:"inputLossBehavior"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.59.0/docs/resources/medialive_channel#output_locking_mode MedialiveChannel#output_locking_mode}.
	OutputLockingMode *string `field:"optional" json:"outputLockingMode" yaml:"outputLockingMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.59.0/docs/resources/medialive_channel#output_timing_source MedialiveChannel#output_timing_source}.
	OutputTimingSource *string `field:"optional" json:"outputTimingSource" yaml:"outputTimingSource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.59.0/docs/resources/medialive_channel#support_low_framerate_inputs MedialiveChannel#support_low_framerate_inputs}.
	SupportLowFramerateInputs *string `field:"optional" json:"supportLowFramerateInputs" yaml:"supportLowFramerateInputs"`
}

