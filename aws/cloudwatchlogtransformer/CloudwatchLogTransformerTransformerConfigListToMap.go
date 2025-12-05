// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudwatchlogtransformer


type CloudwatchLogTransformerTransformerConfigListToMap struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/cloudwatch_log_transformer#key CloudwatchLogTransformer#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/cloudwatch_log_transformer#source CloudwatchLogTransformer#source}.
	Source *string `field:"required" json:"source" yaml:"source"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/cloudwatch_log_transformer#flatten CloudwatchLogTransformer#flatten}.
	Flatten interface{} `field:"optional" json:"flatten" yaml:"flatten"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/cloudwatch_log_transformer#flattened_element CloudwatchLogTransformer#flattened_element}.
	FlattenedElement *string `field:"optional" json:"flattenedElement" yaml:"flattenedElement"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/cloudwatch_log_transformer#target CloudwatchLogTransformer#target}.
	Target *string `field:"optional" json:"target" yaml:"target"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/cloudwatch_log_transformer#value_key CloudwatchLogTransformer#value_key}.
	ValueKey *string `field:"optional" json:"valueKey" yaml:"valueKey"`
}

