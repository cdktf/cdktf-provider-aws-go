package iottopicrule


type IotTopicRuleErrorActionIotEvents struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/resources/iot_topic_rule#input_name IotTopicRule#input_name}.
	InputName *string `field:"required" json:"inputName" yaml:"inputName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/resources/iot_topic_rule#role_arn IotTopicRule#role_arn}.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/resources/iot_topic_rule#batch_mode IotTopicRule#batch_mode}.
	BatchMode interface{} `field:"optional" json:"batchMode" yaml:"batchMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/resources/iot_topic_rule#message_id IotTopicRule#message_id}.
	MessageId *string `field:"optional" json:"messageId" yaml:"messageId"`
}

