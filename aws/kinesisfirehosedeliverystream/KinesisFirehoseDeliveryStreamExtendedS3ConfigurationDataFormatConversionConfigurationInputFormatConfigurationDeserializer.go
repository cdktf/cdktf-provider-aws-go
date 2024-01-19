// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kinesisfirehosedeliverystream


type KinesisFirehoseDeliveryStreamExtendedS3ConfigurationDataFormatConversionConfigurationInputFormatConfigurationDeserializer struct {
	// hive_json_ser_de block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/kinesis_firehose_delivery_stream#hive_json_ser_de KinesisFirehoseDeliveryStream#hive_json_ser_de}
	HiveJsonSerDe *KinesisFirehoseDeliveryStreamExtendedS3ConfigurationDataFormatConversionConfigurationInputFormatConfigurationDeserializerHiveJsonSerDe `field:"optional" json:"hiveJsonSerDe" yaml:"hiveJsonSerDe"`
	// open_x_json_ser_de block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/kinesis_firehose_delivery_stream#open_x_json_ser_de KinesisFirehoseDeliveryStream#open_x_json_ser_de}
	OpenXJsonSerDe *KinesisFirehoseDeliveryStreamExtendedS3ConfigurationDataFormatConversionConfigurationInputFormatConfigurationDeserializerOpenXJsonSerDe `field:"optional" json:"openXJsonSerDe" yaml:"openXJsonSerDe"`
}

