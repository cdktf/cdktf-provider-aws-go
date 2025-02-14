// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rekognitionstreamprocessor


type RekognitionStreamProcessorOutputS3Destination struct {
	// The name of the Amazon S3 bucket you want to associate with the streaming video project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/rekognition_stream_processor#bucket RekognitionStreamProcessor#bucket}
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	// The prefix value of the location within the bucket that you want the information to be published to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/rekognition_stream_processor#key_prefix RekognitionStreamProcessor#key_prefix}
	KeyPrefix *string `field:"optional" json:"keyPrefix" yaml:"keyPrefix"`
}

