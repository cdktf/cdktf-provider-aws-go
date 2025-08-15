// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kinesisstream


type KinesisStreamStreamModeDetails struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/resources/kinesis_stream#stream_mode KinesisStream#stream_mode}.
	StreamMode *string `field:"required" json:"streamMode" yaml:"streamMode"`
}

