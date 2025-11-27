// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iottopicrule


type IotTopicRuleErrorActionTimestreamDimension struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/iot_topic_rule#name IotTopicRule#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/iot_topic_rule#value IotTopicRule#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

