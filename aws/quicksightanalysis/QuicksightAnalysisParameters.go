// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package quicksightanalysis


type QuicksightAnalysisParameters struct {
	// date_time_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/quicksight_analysis#date_time_parameters QuicksightAnalysis#date_time_parameters}
	DateTimeParameters interface{} `field:"optional" json:"dateTimeParameters" yaml:"dateTimeParameters"`
	// decimal_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/quicksight_analysis#decimal_parameters QuicksightAnalysis#decimal_parameters}
	DecimalParameters interface{} `field:"optional" json:"decimalParameters" yaml:"decimalParameters"`
	// integer_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/quicksight_analysis#integer_parameters QuicksightAnalysis#integer_parameters}
	IntegerParameters interface{} `field:"optional" json:"integerParameters" yaml:"integerParameters"`
	// string_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/quicksight_analysis#string_parameters QuicksightAnalysis#string_parameters}
	StringParameters interface{} `field:"optional" json:"stringParameters" yaml:"stringParameters"`
}

