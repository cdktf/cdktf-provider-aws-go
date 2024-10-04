// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudwatcheventtarget


type CloudwatchEventTargetInputTransformer struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.70.0/docs/resources/cloudwatch_event_target#input_template CloudwatchEventTarget#input_template}.
	InputTemplate *string `field:"required" json:"inputTemplate" yaml:"inputTemplate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.70.0/docs/resources/cloudwatch_event_target#input_paths CloudwatchEventTarget#input_paths}.
	InputPaths *map[string]*string `field:"optional" json:"inputPaths" yaml:"inputPaths"`
}

