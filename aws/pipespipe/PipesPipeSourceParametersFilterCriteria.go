// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pipespipe


type PipesPipeSourceParametersFilterCriteria struct {
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/resources/pipes_pipe#filter PipesPipe#filter}
	Filter interface{} `field:"optional" json:"filter" yaml:"filter"`
}

