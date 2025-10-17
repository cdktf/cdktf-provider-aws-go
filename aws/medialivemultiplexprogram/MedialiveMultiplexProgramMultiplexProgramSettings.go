// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package medialivemultiplexprogram


type MedialiveMultiplexProgramMultiplexProgramSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.17.0/docs/resources/medialive_multiplex_program#preferred_channel_pipeline MedialiveMultiplexProgram#preferred_channel_pipeline}.
	PreferredChannelPipeline *string `field:"required" json:"preferredChannelPipeline" yaml:"preferredChannelPipeline"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.17.0/docs/resources/medialive_multiplex_program#program_number MedialiveMultiplexProgram#program_number}.
	ProgramNumber *float64 `field:"required" json:"programNumber" yaml:"programNumber"`
	// service_descriptor block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.17.0/docs/resources/medialive_multiplex_program#service_descriptor MedialiveMultiplexProgram#service_descriptor}
	ServiceDescriptor interface{} `field:"optional" json:"serviceDescriptor" yaml:"serviceDescriptor"`
	// video_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.17.0/docs/resources/medialive_multiplex_program#video_settings MedialiveMultiplexProgram#video_settings}
	VideoSettings interface{} `field:"optional" json:"videoSettings" yaml:"videoSettings"`
}

