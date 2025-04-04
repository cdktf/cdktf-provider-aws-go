// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iottopicrule


type IotTopicRuleTimestream struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.94.0/docs/resources/iot_topic_rule#database_name IotTopicRule#database_name}.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// dimension block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.94.0/docs/resources/iot_topic_rule#dimension IotTopicRule#dimension}
	Dimension interface{} `field:"required" json:"dimension" yaml:"dimension"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.94.0/docs/resources/iot_topic_rule#role_arn IotTopicRule#role_arn}.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.94.0/docs/resources/iot_topic_rule#table_name IotTopicRule#table_name}.
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
	// timestamp block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.94.0/docs/resources/iot_topic_rule#timestamp IotTopicRule#timestamp}
	Timestamp *IotTopicRuleTimestreamTimestamp `field:"optional" json:"timestamp" yaml:"timestamp"`
}

