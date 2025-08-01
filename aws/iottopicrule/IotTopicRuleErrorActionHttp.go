// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iottopicrule


type IotTopicRuleErrorActionHttp struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/iot_topic_rule#url IotTopicRule#url}.
	Url *string `field:"required" json:"url" yaml:"url"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/iot_topic_rule#confirmation_url IotTopicRule#confirmation_url}.
	ConfirmationUrl *string `field:"optional" json:"confirmationUrl" yaml:"confirmationUrl"`
	// http_header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/iot_topic_rule#http_header IotTopicRule#http_header}
	HttpHeader interface{} `field:"optional" json:"httpHeader" yaml:"httpHeader"`
}

