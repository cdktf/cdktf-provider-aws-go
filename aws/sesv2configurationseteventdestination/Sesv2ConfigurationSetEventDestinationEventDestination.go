// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sesv2configurationseteventdestination


type Sesv2ConfigurationSetEventDestinationEventDestination struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.69.0/docs/resources/sesv2_configuration_set_event_destination#matching_event_types Sesv2ConfigurationSetEventDestination#matching_event_types}.
	MatchingEventTypes *[]*string `field:"required" json:"matchingEventTypes" yaml:"matchingEventTypes"`
	// cloud_watch_destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.69.0/docs/resources/sesv2_configuration_set_event_destination#cloud_watch_destination Sesv2ConfigurationSetEventDestination#cloud_watch_destination}
	CloudWatchDestination *Sesv2ConfigurationSetEventDestinationEventDestinationCloudWatchDestination `field:"optional" json:"cloudWatchDestination" yaml:"cloudWatchDestination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.69.0/docs/resources/sesv2_configuration_set_event_destination#enabled Sesv2ConfigurationSetEventDestination#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// event_bridge_destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.69.0/docs/resources/sesv2_configuration_set_event_destination#event_bridge_destination Sesv2ConfigurationSetEventDestination#event_bridge_destination}
	EventBridgeDestination *Sesv2ConfigurationSetEventDestinationEventDestinationEventBridgeDestination `field:"optional" json:"eventBridgeDestination" yaml:"eventBridgeDestination"`
	// kinesis_firehose_destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.69.0/docs/resources/sesv2_configuration_set_event_destination#kinesis_firehose_destination Sesv2ConfigurationSetEventDestination#kinesis_firehose_destination}
	KinesisFirehoseDestination *Sesv2ConfigurationSetEventDestinationEventDestinationKinesisFirehoseDestination `field:"optional" json:"kinesisFirehoseDestination" yaml:"kinesisFirehoseDestination"`
	// pinpoint_destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.69.0/docs/resources/sesv2_configuration_set_event_destination#pinpoint_destination Sesv2ConfigurationSetEventDestination#pinpoint_destination}
	PinpointDestination *Sesv2ConfigurationSetEventDestinationEventDestinationPinpointDestination `field:"optional" json:"pinpointDestination" yaml:"pinpointDestination"`
	// sns_destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.69.0/docs/resources/sesv2_configuration_set_event_destination#sns_destination Sesv2ConfigurationSetEventDestination#sns_destination}
	SnsDestination *Sesv2ConfigurationSetEventDestinationEventDestinationSnsDestination `field:"optional" json:"snsDestination" yaml:"snsDestination"`
}

