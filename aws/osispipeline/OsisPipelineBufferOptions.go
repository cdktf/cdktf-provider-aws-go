// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package osispipeline


type OsisPipelineBufferOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/osis_pipeline#persistent_buffer_enabled OsisPipeline#persistent_buffer_enabled}.
	PersistentBufferEnabled interface{} `field:"required" json:"persistentBufferEnabled" yaml:"persistentBufferEnabled"`
}

