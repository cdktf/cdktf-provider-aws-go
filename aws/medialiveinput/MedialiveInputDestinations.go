// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package medialiveinput


type MedialiveInputDestinations struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/resources/medialive_input#stream_name MedialiveInput#stream_name}.
	StreamName *string `field:"required" json:"streamName" yaml:"streamName"`
}

