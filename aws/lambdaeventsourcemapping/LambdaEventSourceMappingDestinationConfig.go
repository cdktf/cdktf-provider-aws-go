// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lambdaeventsourcemapping


type LambdaEventSourceMappingDestinationConfig struct {
	// on_failure block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/resources/lambda_event_source_mapping#on_failure LambdaEventSourceMapping#on_failure}
	OnFailure *LambdaEventSourceMappingDestinationConfigOnFailure `field:"optional" json:"onFailure" yaml:"onFailure"`
}

