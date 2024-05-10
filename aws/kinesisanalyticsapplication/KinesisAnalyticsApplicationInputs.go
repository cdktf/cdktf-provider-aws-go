// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kinesisanalyticsapplication


type KinesisAnalyticsApplicationInputs struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.49.0/docs/resources/kinesis_analytics_application#name_prefix KinesisAnalyticsApplication#name_prefix}.
	NamePrefix *string `field:"required" json:"namePrefix" yaml:"namePrefix"`
	// schema block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.49.0/docs/resources/kinesis_analytics_application#schema KinesisAnalyticsApplication#schema}
	Schema *KinesisAnalyticsApplicationInputsSchema `field:"required" json:"schema" yaml:"schema"`
	// kinesis_firehose block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.49.0/docs/resources/kinesis_analytics_application#kinesis_firehose KinesisAnalyticsApplication#kinesis_firehose}
	KinesisFirehose *KinesisAnalyticsApplicationInputsKinesisFirehose `field:"optional" json:"kinesisFirehose" yaml:"kinesisFirehose"`
	// kinesis_stream block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.49.0/docs/resources/kinesis_analytics_application#kinesis_stream KinesisAnalyticsApplication#kinesis_stream}
	KinesisStream *KinesisAnalyticsApplicationInputsKinesisStream `field:"optional" json:"kinesisStream" yaml:"kinesisStream"`
	// parallelism block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.49.0/docs/resources/kinesis_analytics_application#parallelism KinesisAnalyticsApplication#parallelism}
	Parallelism *KinesisAnalyticsApplicationInputsParallelism `field:"optional" json:"parallelism" yaml:"parallelism"`
	// processing_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.49.0/docs/resources/kinesis_analytics_application#processing_configuration KinesisAnalyticsApplication#processing_configuration}
	ProcessingConfiguration *KinesisAnalyticsApplicationInputsProcessingConfiguration `field:"optional" json:"processingConfiguration" yaml:"processingConfiguration"`
	// starting_position_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.49.0/docs/resources/kinesis_analytics_application#starting_position_configuration KinesisAnalyticsApplication#starting_position_configuration}
	StartingPositionConfiguration interface{} `field:"optional" json:"startingPositionConfiguration" yaml:"startingPositionConfiguration"`
}

