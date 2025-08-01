// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kinesisfirehosedeliverystream


type KinesisFirehoseDeliveryStreamExtendedS3ConfigurationDataFormatConversionConfigurationOutputFormatConfigurationSerializerOrcSerDe struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/kinesis_firehose_delivery_stream#block_size_bytes KinesisFirehoseDeliveryStream#block_size_bytes}.
	BlockSizeBytes *float64 `field:"optional" json:"blockSizeBytes" yaml:"blockSizeBytes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/kinesis_firehose_delivery_stream#bloom_filter_columns KinesisFirehoseDeliveryStream#bloom_filter_columns}.
	BloomFilterColumns *[]*string `field:"optional" json:"bloomFilterColumns" yaml:"bloomFilterColumns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/kinesis_firehose_delivery_stream#bloom_filter_false_positive_probability KinesisFirehoseDeliveryStream#bloom_filter_false_positive_probability}.
	BloomFilterFalsePositiveProbability *float64 `field:"optional" json:"bloomFilterFalsePositiveProbability" yaml:"bloomFilterFalsePositiveProbability"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/kinesis_firehose_delivery_stream#compression KinesisFirehoseDeliveryStream#compression}.
	Compression *string `field:"optional" json:"compression" yaml:"compression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/kinesis_firehose_delivery_stream#dictionary_key_threshold KinesisFirehoseDeliveryStream#dictionary_key_threshold}.
	DictionaryKeyThreshold *float64 `field:"optional" json:"dictionaryKeyThreshold" yaml:"dictionaryKeyThreshold"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/kinesis_firehose_delivery_stream#enable_padding KinesisFirehoseDeliveryStream#enable_padding}.
	EnablePadding interface{} `field:"optional" json:"enablePadding" yaml:"enablePadding"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/kinesis_firehose_delivery_stream#format_version KinesisFirehoseDeliveryStream#format_version}.
	FormatVersion *string `field:"optional" json:"formatVersion" yaml:"formatVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/kinesis_firehose_delivery_stream#padding_tolerance KinesisFirehoseDeliveryStream#padding_tolerance}.
	PaddingTolerance *float64 `field:"optional" json:"paddingTolerance" yaml:"paddingTolerance"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/kinesis_firehose_delivery_stream#row_index_stride KinesisFirehoseDeliveryStream#row_index_stride}.
	RowIndexStride *float64 `field:"optional" json:"rowIndexStride" yaml:"rowIndexStride"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/kinesis_firehose_delivery_stream#stripe_size_bytes KinesisFirehoseDeliveryStream#stripe_size_bytes}.
	StripeSizeBytes *float64 `field:"optional" json:"stripeSizeBytes" yaml:"stripeSizeBytes"`
}

