// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pipespipe


type PipesPipeTargetParametersKinesisStreamParameters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/pipes_pipe#partition_key PipesPipe#partition_key}.
	PartitionKey *string `field:"required" json:"partitionKey" yaml:"partitionKey"`
}

