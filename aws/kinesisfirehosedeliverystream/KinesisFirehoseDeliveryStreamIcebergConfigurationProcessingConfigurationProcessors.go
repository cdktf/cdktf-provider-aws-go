// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kinesisfirehosedeliverystream


type KinesisFirehoseDeliveryStreamIcebergConfigurationProcessingConfigurationProcessors struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.74.0/docs/resources/kinesis_firehose_delivery_stream#type KinesisFirehoseDeliveryStream#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.74.0/docs/resources/kinesis_firehose_delivery_stream#parameters KinesisFirehoseDeliveryStream#parameters}
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
}

