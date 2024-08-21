// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lambdaeventsourcemapping


type LambdaEventSourceMappingSourceAccessConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.1/docs/resources/lambda_event_source_mapping#type LambdaEventSourceMapping#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.1/docs/resources/lambda_event_source_mapping#uri LambdaEventSourceMapping#uri}.
	Uri *string `field:"required" json:"uri" yaml:"uri"`
}

