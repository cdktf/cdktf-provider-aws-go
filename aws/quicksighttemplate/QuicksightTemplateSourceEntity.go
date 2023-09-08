// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package quicksighttemplate


type QuicksightTemplateSourceEntity struct {
	// source_analysis block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.0/docs/resources/quicksight_template#source_analysis QuicksightTemplate#source_analysis}
	SourceAnalysis *QuicksightTemplateSourceEntitySourceAnalysis `field:"optional" json:"sourceAnalysis" yaml:"sourceAnalysis"`
	// source_template block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.0/docs/resources/quicksight_template#source_template QuicksightTemplate#source_template}
	SourceTemplate *QuicksightTemplateSourceEntitySourceTemplate `field:"optional" json:"sourceTemplate" yaml:"sourceTemplate"`
}

