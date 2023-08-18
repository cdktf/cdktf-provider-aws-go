package kinesisfirehosedeliverystream


type KinesisFirehoseDeliveryStreamKinesisSourceConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/kinesis_firehose_delivery_stream#kinesis_stream_arn KinesisFirehoseDeliveryStream#kinesis_stream_arn}.
	KinesisStreamArn *string `field:"required" json:"kinesisStreamArn" yaml:"kinesisStreamArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/kinesis_firehose_delivery_stream#role_arn KinesisFirehoseDeliveryStream#role_arn}.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
}

