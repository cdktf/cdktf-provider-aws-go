package iottopicrule


type IotTopicRuleDynamodb struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/iot_topic_rule#hash_key_field IotTopicRule#hash_key_field}.
	HashKeyField *string `field:"required" json:"hashKeyField" yaml:"hashKeyField"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/iot_topic_rule#hash_key_value IotTopicRule#hash_key_value}.
	HashKeyValue *string `field:"required" json:"hashKeyValue" yaml:"hashKeyValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/iot_topic_rule#role_arn IotTopicRule#role_arn}.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/iot_topic_rule#table_name IotTopicRule#table_name}.
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/iot_topic_rule#hash_key_type IotTopicRule#hash_key_type}.
	HashKeyType *string `field:"optional" json:"hashKeyType" yaml:"hashKeyType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/iot_topic_rule#operation IotTopicRule#operation}.
	Operation *string `field:"optional" json:"operation" yaml:"operation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/iot_topic_rule#payload_field IotTopicRule#payload_field}.
	PayloadField *string `field:"optional" json:"payloadField" yaml:"payloadField"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/iot_topic_rule#range_key_field IotTopicRule#range_key_field}.
	RangeKeyField *string `field:"optional" json:"rangeKeyField" yaml:"rangeKeyField"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/iot_topic_rule#range_key_type IotTopicRule#range_key_type}.
	RangeKeyType *string `field:"optional" json:"rangeKeyType" yaml:"rangeKeyType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/iot_topic_rule#range_key_value IotTopicRule#range_key_value}.
	RangeKeyValue *string `field:"optional" json:"rangeKeyValue" yaml:"rangeKeyValue"`
}

