// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudwatcheventbus


type CloudwatchEventBusLogConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.11.0/docs/resources/cloudwatch_event_bus#include_detail CloudwatchEventBus#include_detail}.
	IncludeDetail *string `field:"optional" json:"includeDetail" yaml:"includeDetail"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.11.0/docs/resources/cloudwatch_event_bus#level CloudwatchEventBus#level}.
	Level *string `field:"optional" json:"level" yaml:"level"`
}

