package mskconnectconnector


type MskconnectConnectorLogDeliveryWorkerLogDelivery struct {
	// cloudwatch_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/mskconnect_connector#cloudwatch_logs MskconnectConnector#cloudwatch_logs}
	CloudwatchLogs *MskconnectConnectorLogDeliveryWorkerLogDeliveryCloudwatchLogs `field:"optional" json:"cloudwatchLogs" yaml:"cloudwatchLogs"`
	// firehose block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/mskconnect_connector#firehose MskconnectConnector#firehose}
	Firehose *MskconnectConnectorLogDeliveryWorkerLogDeliveryFirehose `field:"optional" json:"firehose" yaml:"firehose"`
	// s3 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/mskconnect_connector#s3 MskconnectConnector#s3}
	S3 *MskconnectConnectorLogDeliveryWorkerLogDeliveryS3 `field:"optional" json:"s3" yaml:"s3"`
}

