// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudwatcheventtarget


type CloudwatchEventTargetSagemakerPipelineTargetPipelineParameterListStruct struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.98.0/docs/resources/cloudwatch_event_target#name CloudwatchEventTarget#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.98.0/docs/resources/cloudwatch_event_target#value CloudwatchEventTarget#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

