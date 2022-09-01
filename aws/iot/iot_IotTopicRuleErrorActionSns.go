package iot


type IotTopicRuleErrorActionSns struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule#role_arn IotTopicRule#role_arn}.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule#target_arn IotTopicRule#target_arn}.
	TargetArn *string `field:"required" json:"targetArn" yaml:"targetArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule#message_format IotTopicRule#message_format}.
	MessageFormat *string `field:"optional" json:"messageFormat" yaml:"messageFormat"`
}

