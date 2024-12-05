// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package medialivechannel


type MedialiveChannelCdiInputSpecification struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.80.0/docs/resources/medialive_channel#resolution MedialiveChannel#resolution}.
	Resolution *string `field:"required" json:"resolution" yaml:"resolution"`
}

