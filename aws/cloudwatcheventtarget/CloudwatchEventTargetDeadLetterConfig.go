// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudwatcheventtarget


type CloudwatchEventTargetDeadLetterConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.49.0/docs/resources/cloudwatch_event_target#arn CloudwatchEventTarget#arn}.
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
}

