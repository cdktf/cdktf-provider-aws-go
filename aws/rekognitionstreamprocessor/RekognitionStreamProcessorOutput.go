// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rekognitionstreamprocessor


type RekognitionStreamProcessorOutput struct {
	// kinesis_data_stream block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/resources/rekognition_stream_processor#kinesis_data_stream RekognitionStreamProcessor#kinesis_data_stream}
	KinesisDataStream interface{} `field:"optional" json:"kinesisDataStream" yaml:"kinesisDataStream"`
	// s3_destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/resources/rekognition_stream_processor#s3_destination RekognitionStreamProcessor#s3_destination}
	S3Destination interface{} `field:"optional" json:"s3Destination" yaml:"s3Destination"`
}

