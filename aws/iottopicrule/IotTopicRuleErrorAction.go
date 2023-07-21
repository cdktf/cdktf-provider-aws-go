package iottopicrule


type IotTopicRuleErrorAction struct {
	// cloudwatch_alarm block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/iot_topic_rule#cloudwatch_alarm IotTopicRule#cloudwatch_alarm}
	CloudwatchAlarm *IotTopicRuleErrorActionCloudwatchAlarm `field:"optional" json:"cloudwatchAlarm" yaml:"cloudwatchAlarm"`
	// cloudwatch_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/iot_topic_rule#cloudwatch_logs IotTopicRule#cloudwatch_logs}
	CloudwatchLogs *IotTopicRuleErrorActionCloudwatchLogs `field:"optional" json:"cloudwatchLogs" yaml:"cloudwatchLogs"`
	// cloudwatch_metric block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/iot_topic_rule#cloudwatch_metric IotTopicRule#cloudwatch_metric}
	CloudwatchMetric *IotTopicRuleErrorActionCloudwatchMetric `field:"optional" json:"cloudwatchMetric" yaml:"cloudwatchMetric"`
	// dynamodb block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/iot_topic_rule#dynamodb IotTopicRule#dynamodb}
	Dynamodb *IotTopicRuleErrorActionDynamodb `field:"optional" json:"dynamodb" yaml:"dynamodb"`
	// dynamodbv2 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/iot_topic_rule#dynamodbv2 IotTopicRule#dynamodbv2}
	Dynamodbv2 *IotTopicRuleErrorActionDynamodbv2 `field:"optional" json:"dynamodbv2" yaml:"dynamodbv2"`
	// elasticsearch block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/iot_topic_rule#elasticsearch IotTopicRule#elasticsearch}
	Elasticsearch *IotTopicRuleErrorActionElasticsearch `field:"optional" json:"elasticsearch" yaml:"elasticsearch"`
	// firehose block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/iot_topic_rule#firehose IotTopicRule#firehose}
	Firehose *IotTopicRuleErrorActionFirehose `field:"optional" json:"firehose" yaml:"firehose"`
	// http block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/iot_topic_rule#http IotTopicRule#http}
	Http *IotTopicRuleErrorActionHttp `field:"optional" json:"http" yaml:"http"`
	// iot_analytics block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/iot_topic_rule#iot_analytics IotTopicRule#iot_analytics}
	IotAnalytics *IotTopicRuleErrorActionIotAnalytics `field:"optional" json:"iotAnalytics" yaml:"iotAnalytics"`
	// iot_events block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/iot_topic_rule#iot_events IotTopicRule#iot_events}
	IotEvents *IotTopicRuleErrorActionIotEvents `field:"optional" json:"iotEvents" yaml:"iotEvents"`
	// kafka block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/iot_topic_rule#kafka IotTopicRule#kafka}
	Kafka *IotTopicRuleErrorActionKafka `field:"optional" json:"kafka" yaml:"kafka"`
	// kinesis block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/iot_topic_rule#kinesis IotTopicRule#kinesis}
	Kinesis *IotTopicRuleErrorActionKinesis `field:"optional" json:"kinesis" yaml:"kinesis"`
	// lambda block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/iot_topic_rule#lambda IotTopicRule#lambda}
	Lambda *IotTopicRuleErrorActionLambda `field:"optional" json:"lambda" yaml:"lambda"`
	// republish block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/iot_topic_rule#republish IotTopicRule#republish}
	Republish *IotTopicRuleErrorActionRepublish `field:"optional" json:"republish" yaml:"republish"`
	// s3 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/iot_topic_rule#s3 IotTopicRule#s3}
	S3 *IotTopicRuleErrorActionS3 `field:"optional" json:"s3" yaml:"s3"`
	// sns block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/iot_topic_rule#sns IotTopicRule#sns}
	Sns *IotTopicRuleErrorActionSns `field:"optional" json:"sns" yaml:"sns"`
	// sqs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/iot_topic_rule#sqs IotTopicRule#sqs}
	Sqs *IotTopicRuleErrorActionSqs `field:"optional" json:"sqs" yaml:"sqs"`
	// step_functions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/iot_topic_rule#step_functions IotTopicRule#step_functions}
	StepFunctions *IotTopicRuleErrorActionStepFunctions `field:"optional" json:"stepFunctions" yaml:"stepFunctions"`
	// timestream block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/iot_topic_rule#timestream IotTopicRule#timestream}
	Timestream *IotTopicRuleErrorActionTimestream `field:"optional" json:"timestream" yaml:"timestream"`
}

