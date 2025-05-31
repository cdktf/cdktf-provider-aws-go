// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iottopicrule


type IotTopicRuleKafka struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/iot_topic_rule#client_properties IotTopicRule#client_properties}.
	ClientProperties *map[string]*string `field:"required" json:"clientProperties" yaml:"clientProperties"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/iot_topic_rule#destination_arn IotTopicRule#destination_arn}.
	DestinationArn *string `field:"required" json:"destinationArn" yaml:"destinationArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/iot_topic_rule#topic IotTopicRule#topic}.
	Topic *string `field:"required" json:"topic" yaml:"topic"`
	// header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/iot_topic_rule#header IotTopicRule#header}
	Header interface{} `field:"optional" json:"header" yaml:"header"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/iot_topic_rule#key IotTopicRule#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/iot_topic_rule#partition IotTopicRule#partition}.
	Partition *string `field:"optional" json:"partition" yaml:"partition"`
}

