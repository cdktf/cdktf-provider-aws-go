// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rekognitionstreamprocessor


type RekognitionStreamProcessorSettingsFaceSearch struct {
	// The ID of a collection that contains faces that you want to search for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.20.0/docs/resources/rekognition_stream_processor#collection_id RekognitionStreamProcessor#collection_id}
	CollectionId *string `field:"required" json:"collectionId" yaml:"collectionId"`
	// Minimum face match confidence score that must be met to return a result for a recognized face.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.20.0/docs/resources/rekognition_stream_processor#face_match_threshold RekognitionStreamProcessor#face_match_threshold}
	FaceMatchThreshold *float64 `field:"optional" json:"faceMatchThreshold" yaml:"faceMatchThreshold"`
}

