// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package medialivechannel


type MedialiveChannelInputAttachmentsInputSettingsAudioSelectorSelectorSettings struct {
	// audio_hls_rendition_selection block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.86.1/docs/resources/medialive_channel#audio_hls_rendition_selection MedialiveChannel#audio_hls_rendition_selection}
	AudioHlsRenditionSelection *MedialiveChannelInputAttachmentsInputSettingsAudioSelectorSelectorSettingsAudioHlsRenditionSelection `field:"optional" json:"audioHlsRenditionSelection" yaml:"audioHlsRenditionSelection"`
	// audio_language_selection block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.86.1/docs/resources/medialive_channel#audio_language_selection MedialiveChannel#audio_language_selection}
	AudioLanguageSelection *MedialiveChannelInputAttachmentsInputSettingsAudioSelectorSelectorSettingsAudioLanguageSelection `field:"optional" json:"audioLanguageSelection" yaml:"audioLanguageSelection"`
	// audio_pid_selection block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.86.1/docs/resources/medialive_channel#audio_pid_selection MedialiveChannel#audio_pid_selection}
	AudioPidSelection *MedialiveChannelInputAttachmentsInputSettingsAudioSelectorSelectorSettingsAudioPidSelection `field:"optional" json:"audioPidSelection" yaml:"audioPidSelection"`
	// audio_track_selection block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.86.1/docs/resources/medialive_channel#audio_track_selection MedialiveChannel#audio_track_selection}
	AudioTrackSelection *MedialiveChannelInputAttachmentsInputSettingsAudioSelectorSelectorSettingsAudioTrackSelection `field:"optional" json:"audioTrackSelection" yaml:"audioTrackSelection"`
}

