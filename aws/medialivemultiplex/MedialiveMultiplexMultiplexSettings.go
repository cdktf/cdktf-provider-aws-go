// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package medialivemultiplex


type MedialiveMultiplexMultiplexSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/medialive_multiplex#transport_stream_bitrate MedialiveMultiplex#transport_stream_bitrate}.
	TransportStreamBitrate *float64 `field:"required" json:"transportStreamBitrate" yaml:"transportStreamBitrate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/medialive_multiplex#transport_stream_id MedialiveMultiplex#transport_stream_id}.
	TransportStreamId *float64 `field:"required" json:"transportStreamId" yaml:"transportStreamId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/medialive_multiplex#maximum_video_buffer_delay_milliseconds MedialiveMultiplex#maximum_video_buffer_delay_milliseconds}.
	MaximumVideoBufferDelayMilliseconds *float64 `field:"optional" json:"maximumVideoBufferDelayMilliseconds" yaml:"maximumVideoBufferDelayMilliseconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/medialive_multiplex#transport_stream_reserved_bitrate MedialiveMultiplex#transport_stream_reserved_bitrate}.
	TransportStreamReservedBitrate *float64 `field:"optional" json:"transportStreamReservedBitrate" yaml:"transportStreamReservedBitrate"`
}

