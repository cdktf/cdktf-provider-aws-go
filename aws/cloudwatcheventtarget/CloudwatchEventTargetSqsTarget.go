// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudwatcheventtarget


type CloudwatchEventTargetSqsTarget struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/cloudwatch_event_target#message_group_id CloudwatchEventTarget#message_group_id}.
	MessageGroupId *string `field:"optional" json:"messageGroupId" yaml:"messageGroupId"`
}

