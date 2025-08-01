// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package quicksighttemplate


type QuicksightTemplateSourceEntitySourceAnalysisDataSetReferences struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/quicksight_template#data_set_arn QuicksightTemplate#data_set_arn}.
	DataSetArn *string `field:"required" json:"dataSetArn" yaml:"dataSetArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/quicksight_template#data_set_placeholder QuicksightTemplate#data_set_placeholder}.
	DataSetPlaceholder *string `field:"required" json:"dataSetPlaceholder" yaml:"dataSetPlaceholder"`
}

