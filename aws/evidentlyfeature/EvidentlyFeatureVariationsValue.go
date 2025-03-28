// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package evidentlyfeature


type EvidentlyFeatureVariationsValue struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.93.0/docs/resources/evidently_feature#bool_value EvidentlyFeature#bool_value}.
	BoolValue *string `field:"optional" json:"boolValue" yaml:"boolValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.93.0/docs/resources/evidently_feature#double_value EvidentlyFeature#double_value}.
	DoubleValue *string `field:"optional" json:"doubleValue" yaml:"doubleValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.93.0/docs/resources/evidently_feature#long_value EvidentlyFeature#long_value}.
	LongValue *string `field:"optional" json:"longValue" yaml:"longValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.93.0/docs/resources/evidently_feature#string_value EvidentlyFeature#string_value}.
	StringValue *string `field:"optional" json:"stringValue" yaml:"stringValue"`
}

