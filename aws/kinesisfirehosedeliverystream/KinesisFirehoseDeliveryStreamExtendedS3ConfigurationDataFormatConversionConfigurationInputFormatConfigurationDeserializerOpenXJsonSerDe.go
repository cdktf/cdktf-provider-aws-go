// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kinesisfirehosedeliverystream


type KinesisFirehoseDeliveryStreamExtendedS3ConfigurationDataFormatConversionConfigurationInputFormatConfigurationDeserializerOpenXJsonSerDe struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.62.0/docs/resources/kinesis_firehose_delivery_stream#case_insensitive KinesisFirehoseDeliveryStream#case_insensitive}.
	CaseInsensitive interface{} `field:"optional" json:"caseInsensitive" yaml:"caseInsensitive"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.62.0/docs/resources/kinesis_firehose_delivery_stream#column_to_json_key_mappings KinesisFirehoseDeliveryStream#column_to_json_key_mappings}.
	ColumnToJsonKeyMappings *map[string]*string `field:"optional" json:"columnToJsonKeyMappings" yaml:"columnToJsonKeyMappings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.62.0/docs/resources/kinesis_firehose_delivery_stream#convert_dots_in_json_keys_to_underscores KinesisFirehoseDeliveryStream#convert_dots_in_json_keys_to_underscores}.
	ConvertDotsInJsonKeysToUnderscores interface{} `field:"optional" json:"convertDotsInJsonKeysToUnderscores" yaml:"convertDotsInJsonKeysToUnderscores"`
}

