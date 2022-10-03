package kinesisfirehosedeliverystream


type KinesisFirehoseDeliveryStreamHttpEndpointConfigurationRequestConfigurationCommonAttributes struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/kinesis_firehose_delivery_stream#name KinesisFirehoseDeliveryStream#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/kinesis_firehose_delivery_stream#value KinesisFirehoseDeliveryStream#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

