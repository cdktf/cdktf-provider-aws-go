package lambdaeventsourcemapping


type LambdaEventSourceMappingSelfManagedKafkaEventSourceConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/lambda_event_source_mapping#consumer_group_id LambdaEventSourceMapping#consumer_group_id}.
	ConsumerGroupId *string `field:"optional" json:"consumerGroupId" yaml:"consumerGroupId"`
}

