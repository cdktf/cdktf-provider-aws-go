// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package quicksightanalysis


type QuicksightAnalysisSourceEntitySourceTemplateDataSetReferences struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.80.0/docs/resources/quicksight_analysis#data_set_arn QuicksightAnalysis#data_set_arn}.
	DataSetArn *string `field:"required" json:"dataSetArn" yaml:"dataSetArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.80.0/docs/resources/quicksight_analysis#data_set_placeholder QuicksightAnalysis#data_set_placeholder}.
	DataSetPlaceholder *string `field:"required" json:"dataSetPlaceholder" yaml:"dataSetPlaceholder"`
}

