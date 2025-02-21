// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rekognitionstreamprocessor


type RekognitionStreamProcessorRegionsOfInterestBoundingBox struct {
	// Height of the bounding box as a ratio of the overall image height.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.88.0/docs/resources/rekognition_stream_processor#height RekognitionStreamProcessor#height}
	Height *float64 `field:"optional" json:"height" yaml:"height"`
	// Left coordinate of the bounding box as a ratio of overall image width.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.88.0/docs/resources/rekognition_stream_processor#left RekognitionStreamProcessor#left}
	Left *float64 `field:"optional" json:"left" yaml:"left"`
	// Top coordinate of the bounding box as a ratio of overall image height.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.88.0/docs/resources/rekognition_stream_processor#top RekognitionStreamProcessor#top}
	Top *float64 `field:"optional" json:"top" yaml:"top"`
	// Width of the bounding box as a ratio of the overall image width.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.88.0/docs/resources/rekognition_stream_processor#width RekognitionStreamProcessor#width}
	Width *float64 `field:"optional" json:"width" yaml:"width"`
}

