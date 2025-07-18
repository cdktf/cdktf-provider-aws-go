// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iottopicrule


type IotTopicRuleErrorActionHttpHttpHeader struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/iot_topic_rule#key IotTopicRule#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/iot_topic_rule#value IotTopicRule#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

