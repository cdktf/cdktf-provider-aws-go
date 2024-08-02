// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pipespipe


type PipesPipeSourceParameters struct {
	// activemq_broker_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.61.0/docs/resources/pipes_pipe#activemq_broker_parameters PipesPipe#activemq_broker_parameters}
	ActivemqBrokerParameters *PipesPipeSourceParametersActivemqBrokerParameters `field:"optional" json:"activemqBrokerParameters" yaml:"activemqBrokerParameters"`
	// dynamodb_stream_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.61.0/docs/resources/pipes_pipe#dynamodb_stream_parameters PipesPipe#dynamodb_stream_parameters}
	DynamodbStreamParameters *PipesPipeSourceParametersDynamodbStreamParameters `field:"optional" json:"dynamodbStreamParameters" yaml:"dynamodbStreamParameters"`
	// filter_criteria block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.61.0/docs/resources/pipes_pipe#filter_criteria PipesPipe#filter_criteria}
	FilterCriteria *PipesPipeSourceParametersFilterCriteria `field:"optional" json:"filterCriteria" yaml:"filterCriteria"`
	// kinesis_stream_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.61.0/docs/resources/pipes_pipe#kinesis_stream_parameters PipesPipe#kinesis_stream_parameters}
	KinesisStreamParameters *PipesPipeSourceParametersKinesisStreamParameters `field:"optional" json:"kinesisStreamParameters" yaml:"kinesisStreamParameters"`
	// managed_streaming_kafka_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.61.0/docs/resources/pipes_pipe#managed_streaming_kafka_parameters PipesPipe#managed_streaming_kafka_parameters}
	ManagedStreamingKafkaParameters *PipesPipeSourceParametersManagedStreamingKafkaParameters `field:"optional" json:"managedStreamingKafkaParameters" yaml:"managedStreamingKafkaParameters"`
	// rabbitmq_broker_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.61.0/docs/resources/pipes_pipe#rabbitmq_broker_parameters PipesPipe#rabbitmq_broker_parameters}
	RabbitmqBrokerParameters *PipesPipeSourceParametersRabbitmqBrokerParameters `field:"optional" json:"rabbitmqBrokerParameters" yaml:"rabbitmqBrokerParameters"`
	// self_managed_kafka_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.61.0/docs/resources/pipes_pipe#self_managed_kafka_parameters PipesPipe#self_managed_kafka_parameters}
	SelfManagedKafkaParameters *PipesPipeSourceParametersSelfManagedKafkaParameters `field:"optional" json:"selfManagedKafkaParameters" yaml:"selfManagedKafkaParameters"`
	// sqs_queue_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.61.0/docs/resources/pipes_pipe#sqs_queue_parameters PipesPipe#sqs_queue_parameters}
	SqsQueueParameters *PipesPipeSourceParametersSqsQueueParameters `field:"optional" json:"sqsQueueParameters" yaml:"sqsQueueParameters"`
}

