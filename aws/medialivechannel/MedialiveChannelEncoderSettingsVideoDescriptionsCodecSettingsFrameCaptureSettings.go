// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package medialivechannel


type MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsFrameCaptureSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.53.0/docs/resources/medialive_channel#capture_interval MedialiveChannel#capture_interval}.
	CaptureInterval *float64 `field:"optional" json:"captureInterval" yaml:"captureInterval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.53.0/docs/resources/medialive_channel#capture_interval_units MedialiveChannel#capture_interval_units}.
	CaptureIntervalUnits *string `field:"optional" json:"captureIntervalUnits" yaml:"captureIntervalUnits"`
}

