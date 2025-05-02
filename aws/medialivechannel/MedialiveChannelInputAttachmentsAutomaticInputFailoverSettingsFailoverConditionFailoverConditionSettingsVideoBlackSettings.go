// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package medialivechannel


type MedialiveChannelInputAttachmentsAutomaticInputFailoverSettingsFailoverConditionFailoverConditionSettingsVideoBlackSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.97.0/docs/resources/medialive_channel#black_detect_threshold MedialiveChannel#black_detect_threshold}.
	BlackDetectThreshold *float64 `field:"optional" json:"blackDetectThreshold" yaml:"blackDetectThreshold"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.97.0/docs/resources/medialive_channel#video_black_threshold_msec MedialiveChannel#video_black_threshold_msec}.
	VideoBlackThresholdMsec *float64 `field:"optional" json:"videoBlackThresholdMsec" yaml:"videoBlackThresholdMsec"`
}

