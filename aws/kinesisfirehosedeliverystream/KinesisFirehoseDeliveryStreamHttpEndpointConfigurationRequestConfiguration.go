// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kinesisfirehosedeliverystream


type KinesisFirehoseDeliveryStreamHttpEndpointConfigurationRequestConfiguration struct {
	// common_attributes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/kinesis_firehose_delivery_stream#common_attributes KinesisFirehoseDeliveryStream#common_attributes}
	CommonAttributes interface{} `field:"optional" json:"commonAttributes" yaml:"commonAttributes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/kinesis_firehose_delivery_stream#content_encoding KinesisFirehoseDeliveryStream#content_encoding}.
	ContentEncoding *string `field:"optional" json:"contentEncoding" yaml:"contentEncoding"`
}

