// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rekognitionstreamprocessor


type RekognitionStreamProcessorRegionsOfInterestPolygon struct {
	// The value of the X coordinate for a point on a Polygon.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.93.0/docs/resources/rekognition_stream_processor#x RekognitionStreamProcessor#x}
	X *float64 `field:"optional" json:"x" yaml:"x"`
	// The value of the Y coordinate for a point on a Polygon.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.93.0/docs/resources/rekognition_stream_processor#y RekognitionStreamProcessor#y}
	Y *float64 `field:"optional" json:"y" yaml:"y"`
}

