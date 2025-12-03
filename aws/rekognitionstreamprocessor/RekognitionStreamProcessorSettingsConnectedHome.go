// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rekognitionstreamprocessor


type RekognitionStreamProcessorSettingsConnectedHome struct {
	// Specifies what you want to detect in the video, such as people, packages, or pets.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/resources/rekognition_stream_processor#labels RekognitionStreamProcessor#labels}
	Labels *[]*string `field:"optional" json:"labels" yaml:"labels"`
	// The minimum confidence required to label an object in the video.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/resources/rekognition_stream_processor#min_confidence RekognitionStreamProcessor#min_confidence}
	MinConfidence *float64 `field:"optional" json:"minConfidence" yaml:"minConfidence"`
}

