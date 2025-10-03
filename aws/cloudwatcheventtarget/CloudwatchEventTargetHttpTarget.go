// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudwatcheventtarget


type CloudwatchEventTargetHttpTarget struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/cloudwatch_event_target#header_parameters CloudwatchEventTarget#header_parameters}.
	HeaderParameters *map[string]*string `field:"optional" json:"headerParameters" yaml:"headerParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/cloudwatch_event_target#path_parameter_values CloudwatchEventTarget#path_parameter_values}.
	PathParameterValues *[]*string `field:"optional" json:"pathParameterValues" yaml:"pathParameterValues"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/cloudwatch_event_target#query_string_parameters CloudwatchEventTarget#query_string_parameters}.
	QueryStringParameters *map[string]*string `field:"optional" json:"queryStringParameters" yaml:"queryStringParameters"`
}

