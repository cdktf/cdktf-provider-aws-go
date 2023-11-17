// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sesv2configurationseteventdestination


type Sesv2ConfigurationSetEventDestinationEventDestinationCloudWatchDestinationDimensionConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.26.0/docs/resources/sesv2_configuration_set_event_destination#default_dimension_value Sesv2ConfigurationSetEventDestination#default_dimension_value}.
	DefaultDimensionValue *string `field:"required" json:"defaultDimensionValue" yaml:"defaultDimensionValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.26.0/docs/resources/sesv2_configuration_set_event_destination#dimension_name Sesv2ConfigurationSetEventDestination#dimension_name}.
	DimensionName *string `field:"required" json:"dimensionName" yaml:"dimensionName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.26.0/docs/resources/sesv2_configuration_set_event_destination#dimension_value_source Sesv2ConfigurationSetEventDestination#dimension_value_source}.
	DimensionValueSource *string `field:"required" json:"dimensionValueSource" yaml:"dimensionValueSource"`
}

