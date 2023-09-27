// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudwatcheventtarget


type CloudwatchEventTargetEcsTargetOrderedPlacementStrategy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/cloudwatch_event_target#type CloudwatchEventTarget#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/cloudwatch_event_target#field CloudwatchEventTarget#field}.
	Field *string `field:"optional" json:"field" yaml:"field"`
}

