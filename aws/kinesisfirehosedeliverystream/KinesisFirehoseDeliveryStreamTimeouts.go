package kinesisfirehosedeliverystream


type KinesisFirehoseDeliveryStreamTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/kinesis_firehose_delivery_stream#create KinesisFirehoseDeliveryStream#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/kinesis_firehose_delivery_stream#delete KinesisFirehoseDeliveryStream#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/kinesis_firehose_delivery_stream#update KinesisFirehoseDeliveryStream#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

