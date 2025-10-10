// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudfrontrealtimelogconfig


type CloudfrontRealtimeLogConfigEndpoint struct {
	// kinesis_stream_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/cloudfront_realtime_log_config#kinesis_stream_config CloudfrontRealtimeLogConfig#kinesis_stream_config}
	KinesisStreamConfig *CloudfrontRealtimeLogConfigEndpointKinesisStreamConfig `field:"required" json:"kinesisStreamConfig" yaml:"kinesisStreamConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/cloudfront_realtime_log_config#stream_type CloudfrontRealtimeLogConfig#stream_type}.
	StreamType *string `field:"required" json:"streamType" yaml:"streamType"`
}

