// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudwatcheventbus


type CloudwatchEventBusDeadLetterConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.13.0/docs/resources/cloudwatch_event_bus#arn CloudwatchEventBus#arn}.
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
}

