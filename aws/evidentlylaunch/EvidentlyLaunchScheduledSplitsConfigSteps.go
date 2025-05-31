// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package evidentlylaunch


type EvidentlyLaunchScheduledSplitsConfigSteps struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/evidently_launch#group_weights EvidentlyLaunch#group_weights}.
	GroupWeights *map[string]*float64 `field:"required" json:"groupWeights" yaml:"groupWeights"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/evidently_launch#start_time EvidentlyLaunch#start_time}.
	StartTime *string `field:"required" json:"startTime" yaml:"startTime"`
	// segment_overrides block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/evidently_launch#segment_overrides EvidentlyLaunch#segment_overrides}
	SegmentOverrides interface{} `field:"optional" json:"segmentOverrides" yaml:"segmentOverrides"`
}

