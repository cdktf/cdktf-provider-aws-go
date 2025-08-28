// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package qldbstream


type QldbStreamKinesisConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.11.0/docs/resources/qldb_stream#stream_arn QldbStream#stream_arn}.
	StreamArn *string `field:"required" json:"streamArn" yaml:"streamArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.11.0/docs/resources/qldb_stream#aggregation_enabled QldbStream#aggregation_enabled}.
	AggregationEnabled interface{} `field:"optional" json:"aggregationEnabled" yaml:"aggregationEnabled"`
}

