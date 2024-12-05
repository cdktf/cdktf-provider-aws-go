// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lambdaeventsourcemapping


type LambdaEventSourceMappingFilterCriteria struct {
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.80.0/docs/resources/lambda_event_source_mapping#filter LambdaEventSourceMapping#filter}
	Filter interface{} `field:"optional" json:"filter" yaml:"filter"`
}

