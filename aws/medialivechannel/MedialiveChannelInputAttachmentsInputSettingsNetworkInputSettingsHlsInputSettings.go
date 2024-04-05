// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package medialivechannel


type MedialiveChannelInputAttachmentsInputSettingsNetworkInputSettingsHlsInputSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.44.0/docs/resources/medialive_channel#bandwidth MedialiveChannel#bandwidth}.
	Bandwidth *float64 `field:"optional" json:"bandwidth" yaml:"bandwidth"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.44.0/docs/resources/medialive_channel#buffer_segments MedialiveChannel#buffer_segments}.
	BufferSegments *float64 `field:"optional" json:"bufferSegments" yaml:"bufferSegments"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.44.0/docs/resources/medialive_channel#retries MedialiveChannel#retries}.
	Retries *float64 `field:"optional" json:"retries" yaml:"retries"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.44.0/docs/resources/medialive_channel#retry_interval MedialiveChannel#retry_interval}.
	RetryInterval *float64 `field:"optional" json:"retryInterval" yaml:"retryInterval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.44.0/docs/resources/medialive_channel#scte35_source MedialiveChannel#scte35_source}.
	Scte35Source *string `field:"optional" json:"scte35Source" yaml:"scte35Source"`
}

