package kinesisfirehosedeliverystream


type KinesisFirehoseDeliveryStreamElasticsearchConfigurationProcessingConfigurationProcessors struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/kinesis_firehose_delivery_stream#type KinesisFirehoseDeliveryStream#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// parameters block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/kinesis_firehose_delivery_stream#parameters KinesisFirehoseDeliveryStream#parameters}
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
}

