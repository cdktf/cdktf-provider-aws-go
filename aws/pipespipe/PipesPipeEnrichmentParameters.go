// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pipespipe


type PipesPipeEnrichmentParameters struct {
	// http_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/pipes_pipe#http_parameters PipesPipe#http_parameters}
	HttpParameters *PipesPipeEnrichmentParametersHttpParameters `field:"optional" json:"httpParameters" yaml:"httpParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/pipes_pipe#input_template PipesPipe#input_template}.
	InputTemplate *string `field:"optional" json:"inputTemplate" yaml:"inputTemplate"`
}

