// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iottopicrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IotTopicRuleConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/iot_topic_rule#enabled IotTopicRule#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/iot_topic_rule#name IotTopicRule#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/iot_topic_rule#sql IotTopicRule#sql}.
	Sql *string `field:"required" json:"sql" yaml:"sql"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/iot_topic_rule#sql_version IotTopicRule#sql_version}.
	SqlVersion *string `field:"required" json:"sqlVersion" yaml:"sqlVersion"`
	// cloudwatch_alarm block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/iot_topic_rule#cloudwatch_alarm IotTopicRule#cloudwatch_alarm}
	CloudwatchAlarm interface{} `field:"optional" json:"cloudwatchAlarm" yaml:"cloudwatchAlarm"`
	// cloudwatch_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/iot_topic_rule#cloudwatch_logs IotTopicRule#cloudwatch_logs}
	CloudwatchLogs interface{} `field:"optional" json:"cloudwatchLogs" yaml:"cloudwatchLogs"`
	// cloudwatch_metric block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/iot_topic_rule#cloudwatch_metric IotTopicRule#cloudwatch_metric}
	CloudwatchMetric interface{} `field:"optional" json:"cloudwatchMetric" yaml:"cloudwatchMetric"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/iot_topic_rule#description IotTopicRule#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// dynamodb block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/iot_topic_rule#dynamodb IotTopicRule#dynamodb}
	Dynamodb interface{} `field:"optional" json:"dynamodb" yaml:"dynamodb"`
	// dynamodbv2 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/iot_topic_rule#dynamodbv2 IotTopicRule#dynamodbv2}
	Dynamodbv2 interface{} `field:"optional" json:"dynamodbv2" yaml:"dynamodbv2"`
	// elasticsearch block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/iot_topic_rule#elasticsearch IotTopicRule#elasticsearch}
	Elasticsearch interface{} `field:"optional" json:"elasticsearch" yaml:"elasticsearch"`
	// error_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/iot_topic_rule#error_action IotTopicRule#error_action}
	ErrorAction *IotTopicRuleErrorAction `field:"optional" json:"errorAction" yaml:"errorAction"`
	// firehose block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/iot_topic_rule#firehose IotTopicRule#firehose}
	Firehose interface{} `field:"optional" json:"firehose" yaml:"firehose"`
	// http block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/iot_topic_rule#http IotTopicRule#http}
	Http interface{} `field:"optional" json:"http" yaml:"http"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/iot_topic_rule#id IotTopicRule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// iot_analytics block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/iot_topic_rule#iot_analytics IotTopicRule#iot_analytics}
	IotAnalytics interface{} `field:"optional" json:"iotAnalytics" yaml:"iotAnalytics"`
	// iot_events block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/iot_topic_rule#iot_events IotTopicRule#iot_events}
	IotEvents interface{} `field:"optional" json:"iotEvents" yaml:"iotEvents"`
	// kafka block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/iot_topic_rule#kafka IotTopicRule#kafka}
	Kafka interface{} `field:"optional" json:"kafka" yaml:"kafka"`
	// kinesis block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/iot_topic_rule#kinesis IotTopicRule#kinesis}
	Kinesis interface{} `field:"optional" json:"kinesis" yaml:"kinesis"`
	// lambda block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/iot_topic_rule#lambda IotTopicRule#lambda}
	Lambda interface{} `field:"optional" json:"lambda" yaml:"lambda"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/iot_topic_rule#region IotTopicRule#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// republish block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/iot_topic_rule#republish IotTopicRule#republish}
	Republish interface{} `field:"optional" json:"republish" yaml:"republish"`
	// s3 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/iot_topic_rule#s3 IotTopicRule#s3}
	S3 interface{} `field:"optional" json:"s3" yaml:"s3"`
	// sns block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/iot_topic_rule#sns IotTopicRule#sns}
	Sns interface{} `field:"optional" json:"sns" yaml:"sns"`
	// sqs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/iot_topic_rule#sqs IotTopicRule#sqs}
	Sqs interface{} `field:"optional" json:"sqs" yaml:"sqs"`
	// step_functions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/iot_topic_rule#step_functions IotTopicRule#step_functions}
	StepFunctions interface{} `field:"optional" json:"stepFunctions" yaml:"stepFunctions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/iot_topic_rule#tags IotTopicRule#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/iot_topic_rule#tags_all IotTopicRule#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timestream block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/iot_topic_rule#timestream IotTopicRule#timestream}
	Timestream interface{} `field:"optional" json:"timestream" yaml:"timestream"`
}

