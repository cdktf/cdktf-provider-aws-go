package iot


type IotTopicRuleIotAnalytics struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule#channel_name IotTopicRule#channel_name}.
	ChannelName *string `field:"required" json:"channelName" yaml:"channelName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_topic_rule#role_arn IotTopicRule#role_arn}.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
}

