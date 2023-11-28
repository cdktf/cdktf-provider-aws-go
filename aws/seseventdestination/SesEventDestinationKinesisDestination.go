// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package seseventdestination


type SesEventDestinationKinesisDestination struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.27.0/docs/resources/ses_event_destination#role_arn SesEventDestination#role_arn}.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.27.0/docs/resources/ses_event_destination#stream_arn SesEventDestination#stream_arn}.
	StreamArn *string `field:"required" json:"streamArn" yaml:"streamArn"`
}

