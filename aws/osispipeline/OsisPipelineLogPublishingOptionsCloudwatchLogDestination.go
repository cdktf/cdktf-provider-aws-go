// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package osispipeline


type OsisPipelineLogPublishingOptionsCloudwatchLogDestination struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/osis_pipeline#log_group OsisPipeline#log_group}.
	LogGroup *string `field:"required" json:"logGroup" yaml:"logGroup"`
}

