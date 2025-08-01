// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iottopicrule


type IotTopicRuleErrorActionDynamodbv2PutItem struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/iot_topic_rule#table_name IotTopicRule#table_name}.
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
}

