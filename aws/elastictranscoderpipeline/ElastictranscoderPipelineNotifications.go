// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastictranscoderpipeline


type ElastictranscoderPipelineNotifications struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.46.0/docs/resources/elastictranscoder_pipeline#completed ElastictranscoderPipeline#completed}.
	Completed *string `field:"optional" json:"completed" yaml:"completed"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.46.0/docs/resources/elastictranscoder_pipeline#error ElastictranscoderPipeline#error}.
	Error *string `field:"optional" json:"error" yaml:"error"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.46.0/docs/resources/elastictranscoder_pipeline#progressing ElastictranscoderPipeline#progressing}.
	Progressing *string `field:"optional" json:"progressing" yaml:"progressing"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.46.0/docs/resources/elastictranscoder_pipeline#warning ElastictranscoderPipeline#warning}.
	Warning *string `field:"optional" json:"warning" yaml:"warning"`
}

