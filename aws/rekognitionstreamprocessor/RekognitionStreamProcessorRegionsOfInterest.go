// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rekognitionstreamprocessor


type RekognitionStreamProcessorRegionsOfInterest struct {
	// bounding_box block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/rekognition_stream_processor#bounding_box RekognitionStreamProcessor#bounding_box}
	BoundingBox interface{} `field:"optional" json:"boundingBox" yaml:"boundingBox"`
	// polygon block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/rekognition_stream_processor#polygon RekognitionStreamProcessor#polygon}
	Polygon interface{} `field:"optional" json:"polygon" yaml:"polygon"`
}

