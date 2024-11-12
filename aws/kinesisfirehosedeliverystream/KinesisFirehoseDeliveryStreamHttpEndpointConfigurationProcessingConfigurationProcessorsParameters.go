// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kinesisfirehosedeliverystream


type KinesisFirehoseDeliveryStreamHttpEndpointConfigurationProcessingConfigurationProcessorsParameters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/kinesis_firehose_delivery_stream#parameter_name KinesisFirehoseDeliveryStream#parameter_name}.
	ParameterName *string `field:"required" json:"parameterName" yaml:"parameterName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/kinesis_firehose_delivery_stream#parameter_value KinesisFirehoseDeliveryStream#parameter_value}.
	ParameterValue *string `field:"required" json:"parameterValue" yaml:"parameterValue"`
}

