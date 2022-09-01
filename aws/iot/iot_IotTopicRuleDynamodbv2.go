package iot


type IotTopicRuleDynamodbv2 struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule#role_arn IotTopicRule#role_arn}.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// put_item block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule#put_item IotTopicRule#put_item}
	PutItem *IotTopicRuleDynamodbv2PutItem `field:"optional" json:"putItem" yaml:"putItem"`
}

