// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kinesisanalyticsv2application


type Kinesisanalyticsv2ApplicationApplicationConfigurationSqlApplicationConfigurationOutput struct {
	// destination_schema block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/kinesisanalyticsv2_application#destination_schema Kinesisanalyticsv2Application#destination_schema}
	DestinationSchema *Kinesisanalyticsv2ApplicationApplicationConfigurationSqlApplicationConfigurationOutputDestinationSchema `field:"required" json:"destinationSchema" yaml:"destinationSchema"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/kinesisanalyticsv2_application#name Kinesisanalyticsv2Application#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// kinesis_firehose_output block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/kinesisanalyticsv2_application#kinesis_firehose_output Kinesisanalyticsv2Application#kinesis_firehose_output}
	KinesisFirehoseOutput *Kinesisanalyticsv2ApplicationApplicationConfigurationSqlApplicationConfigurationOutputKinesisFirehoseOutput `field:"optional" json:"kinesisFirehoseOutput" yaml:"kinesisFirehoseOutput"`
	// kinesis_streams_output block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/kinesisanalyticsv2_application#kinesis_streams_output Kinesisanalyticsv2Application#kinesis_streams_output}
	KinesisStreamsOutput *Kinesisanalyticsv2ApplicationApplicationConfigurationSqlApplicationConfigurationOutputKinesisStreamsOutput `field:"optional" json:"kinesisStreamsOutput" yaml:"kinesisStreamsOutput"`
	// lambda_output block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/kinesisanalyticsv2_application#lambda_output Kinesisanalyticsv2Application#lambda_output}
	LambdaOutput *Kinesisanalyticsv2ApplicationApplicationConfigurationSqlApplicationConfigurationOutputLambdaOutput `field:"optional" json:"lambdaOutput" yaml:"lambdaOutput"`
}

