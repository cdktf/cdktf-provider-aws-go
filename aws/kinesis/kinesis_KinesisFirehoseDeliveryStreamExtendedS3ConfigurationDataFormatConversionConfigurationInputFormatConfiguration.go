package kinesis


type KinesisFirehoseDeliveryStreamExtendedS3ConfigurationDataFormatConversionConfigurationInputFormatConfiguration struct {
	// deserializer block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/kinesis_firehose_delivery_stream#deserializer KinesisFirehoseDeliveryStream#deserializer}
	Deserializer *KinesisFirehoseDeliveryStreamExtendedS3ConfigurationDataFormatConversionConfigurationInputFormatConfigurationDeserializer `field:"required" json:"deserializer" yaml:"deserializer"`
}

