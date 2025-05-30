// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudwatcheventtarget


type CloudwatchEventTargetBatchTarget struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/resources/cloudwatch_event_target#job_definition CloudwatchEventTarget#job_definition}.
	JobDefinition *string `field:"required" json:"jobDefinition" yaml:"jobDefinition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/resources/cloudwatch_event_target#job_name CloudwatchEventTarget#job_name}.
	JobName *string `field:"required" json:"jobName" yaml:"jobName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/resources/cloudwatch_event_target#array_size CloudwatchEventTarget#array_size}.
	ArraySize *float64 `field:"optional" json:"arraySize" yaml:"arraySize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/resources/cloudwatch_event_target#job_attempts CloudwatchEventTarget#job_attempts}.
	JobAttempts *float64 `field:"optional" json:"jobAttempts" yaml:"jobAttempts"`
}

