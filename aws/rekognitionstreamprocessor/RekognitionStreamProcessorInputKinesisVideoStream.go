// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rekognitionstreamprocessor


type RekognitionStreamProcessorInputKinesisVideoStream struct {
	// ARN of the Kinesis video stream stream that streams the source video.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.94.0/docs/resources/rekognition_stream_processor#arn RekognitionStreamProcessor#arn}
	Arn *string `field:"required" json:"arn" yaml:"arn"`
}

