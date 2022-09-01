package kinesis


type KinesisFirehoseDeliveryStreamExtendedS3ConfigurationDataFormatConversionConfigurationOutputFormatConfiguration struct {
	// serializer block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/kinesis_firehose_delivery_stream#serializer KinesisFirehoseDeliveryStream#serializer}
	Serializer *KinesisFirehoseDeliveryStreamExtendedS3ConfigurationDataFormatConversionConfigurationOutputFormatConfigurationSerializer `field:"required" json:"serializer" yaml:"serializer"`
}

