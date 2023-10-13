// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sesv2configurationseteventdestination


type Sesv2ConfigurationSetEventDestinationEventDestinationCloudWatchDestination struct {
	// dimension_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.21.0/docs/resources/sesv2_configuration_set_event_destination#dimension_configuration Sesv2ConfigurationSetEventDestination#dimension_configuration}
	DimensionConfiguration interface{} `field:"required" json:"dimensionConfiguration" yaml:"dimensionConfiguration"`
}

