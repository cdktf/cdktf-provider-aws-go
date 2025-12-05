// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudwatchlogtransformer


type CloudwatchLogTransformerTransformerConfigParseKeyValue struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/cloudwatch_log_transformer#destination CloudwatchLogTransformer#destination}.
	Destination *string `field:"optional" json:"destination" yaml:"destination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/cloudwatch_log_transformer#field_delimiter CloudwatchLogTransformer#field_delimiter}.
	FieldDelimiter *string `field:"optional" json:"fieldDelimiter" yaml:"fieldDelimiter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/cloudwatch_log_transformer#key_prefix CloudwatchLogTransformer#key_prefix}.
	KeyPrefix *string `field:"optional" json:"keyPrefix" yaml:"keyPrefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/cloudwatch_log_transformer#key_value_delimiter CloudwatchLogTransformer#key_value_delimiter}.
	KeyValueDelimiter *string `field:"optional" json:"keyValueDelimiter" yaml:"keyValueDelimiter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/cloudwatch_log_transformer#non_match_value CloudwatchLogTransformer#non_match_value}.
	NonMatchValue *string `field:"optional" json:"nonMatchValue" yaml:"nonMatchValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/cloudwatch_log_transformer#overwrite_if_exists CloudwatchLogTransformer#overwrite_if_exists}.
	OverwriteIfExists interface{} `field:"optional" json:"overwriteIfExists" yaml:"overwriteIfExists"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/cloudwatch_log_transformer#source CloudwatchLogTransformer#source}.
	Source *string `field:"optional" json:"source" yaml:"source"`
}

