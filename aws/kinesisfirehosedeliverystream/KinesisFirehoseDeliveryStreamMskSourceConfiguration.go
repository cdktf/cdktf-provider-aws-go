// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kinesisfirehosedeliverystream


type KinesisFirehoseDeliveryStreamMskSourceConfiguration struct {
	// authentication_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.64.0/docs/resources/kinesis_firehose_delivery_stream#authentication_configuration KinesisFirehoseDeliveryStream#authentication_configuration}
	AuthenticationConfiguration *KinesisFirehoseDeliveryStreamMskSourceConfigurationAuthenticationConfiguration `field:"required" json:"authenticationConfiguration" yaml:"authenticationConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.64.0/docs/resources/kinesis_firehose_delivery_stream#msk_cluster_arn KinesisFirehoseDeliveryStream#msk_cluster_arn}.
	MskClusterArn *string `field:"required" json:"mskClusterArn" yaml:"mskClusterArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.64.0/docs/resources/kinesis_firehose_delivery_stream#topic_name KinesisFirehoseDeliveryStream#topic_name}.
	TopicName *string `field:"required" json:"topicName" yaml:"topicName"`
}

