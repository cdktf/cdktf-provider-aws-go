// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package seseventdestination


type SesEventDestinationCloudwatchDestination struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.1/docs/resources/ses_event_destination#default_value SesEventDestination#default_value}.
	DefaultValue *string `field:"required" json:"defaultValue" yaml:"defaultValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.1/docs/resources/ses_event_destination#dimension_name SesEventDestination#dimension_name}.
	DimensionName *string `field:"required" json:"dimensionName" yaml:"dimensionName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.1/docs/resources/ses_event_destination#value_source SesEventDestination#value_source}.
	ValueSource *string `field:"required" json:"valueSource" yaml:"valueSource"`
}

