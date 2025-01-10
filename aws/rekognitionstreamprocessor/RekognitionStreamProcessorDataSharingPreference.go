// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rekognitionstreamprocessor


type RekognitionStreamProcessorDataSharingPreference struct {
	// Do you want to share data with Rekognition to improve model performance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.0/docs/resources/rekognition_stream_processor#opt_in RekognitionStreamProcessor#opt_in}
	OptIn interface{} `field:"required" json:"optIn" yaml:"optIn"`
}

