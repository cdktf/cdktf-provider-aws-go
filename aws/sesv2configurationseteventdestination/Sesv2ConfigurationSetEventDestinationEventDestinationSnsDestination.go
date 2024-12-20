// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sesv2configurationseteventdestination


type Sesv2ConfigurationSetEventDestinationEventDestinationSnsDestination struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/resources/sesv2_configuration_set_event_destination#topic_arn Sesv2ConfigurationSetEventDestination#topic_arn}.
	TopicArn *string `field:"required" json:"topicArn" yaml:"topicArn"`
}

