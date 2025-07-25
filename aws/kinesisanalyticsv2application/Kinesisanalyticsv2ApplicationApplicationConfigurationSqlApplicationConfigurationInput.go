// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kinesisanalyticsv2application


type Kinesisanalyticsv2ApplicationApplicationConfigurationSqlApplicationConfigurationInput struct {
	// input_schema block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/kinesisanalyticsv2_application#input_schema Kinesisanalyticsv2Application#input_schema}
	InputSchema *Kinesisanalyticsv2ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputSchema `field:"required" json:"inputSchema" yaml:"inputSchema"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/kinesisanalyticsv2_application#name_prefix Kinesisanalyticsv2Application#name_prefix}.
	NamePrefix *string `field:"required" json:"namePrefix" yaml:"namePrefix"`
	// input_parallelism block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/kinesisanalyticsv2_application#input_parallelism Kinesisanalyticsv2Application#input_parallelism}
	InputParallelism *Kinesisanalyticsv2ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputParallelism `field:"optional" json:"inputParallelism" yaml:"inputParallelism"`
	// input_processing_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/kinesisanalyticsv2_application#input_processing_configuration Kinesisanalyticsv2Application#input_processing_configuration}
	InputProcessingConfiguration *Kinesisanalyticsv2ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputProcessingConfiguration `field:"optional" json:"inputProcessingConfiguration" yaml:"inputProcessingConfiguration"`
	// input_starting_position_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/kinesisanalyticsv2_application#input_starting_position_configuration Kinesisanalyticsv2Application#input_starting_position_configuration}
	InputStartingPositionConfiguration interface{} `field:"optional" json:"inputStartingPositionConfiguration" yaml:"inputStartingPositionConfiguration"`
	// kinesis_firehose_input block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/kinesisanalyticsv2_application#kinesis_firehose_input Kinesisanalyticsv2Application#kinesis_firehose_input}
	KinesisFirehoseInput *Kinesisanalyticsv2ApplicationApplicationConfigurationSqlApplicationConfigurationInputKinesisFirehoseInput `field:"optional" json:"kinesisFirehoseInput" yaml:"kinesisFirehoseInput"`
	// kinesis_streams_input block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/kinesisanalyticsv2_application#kinesis_streams_input Kinesisanalyticsv2Application#kinesis_streams_input}
	KinesisStreamsInput *Kinesisanalyticsv2ApplicationApplicationConfigurationSqlApplicationConfigurationInputKinesisStreamsInput `field:"optional" json:"kinesisStreamsInput" yaml:"kinesisStreamsInput"`
}

