// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rekognitionstreamprocessor


type RekognitionStreamProcessorSettings struct {
	// connected_home block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.69.0/docs/resources/rekognition_stream_processor#connected_home RekognitionStreamProcessor#connected_home}
	ConnectedHome interface{} `field:"optional" json:"connectedHome" yaml:"connectedHome"`
	// face_search block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.69.0/docs/resources/rekognition_stream_processor#face_search RekognitionStreamProcessor#face_search}
	FaceSearch interface{} `field:"optional" json:"faceSearch" yaml:"faceSearch"`
}

