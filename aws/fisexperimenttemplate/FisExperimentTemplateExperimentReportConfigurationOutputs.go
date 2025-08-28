// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package fisexperimenttemplate


type FisExperimentTemplateExperimentReportConfigurationOutputs struct {
	// s3_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.11.0/docs/resources/fis_experiment_template#s3_configuration FisExperimentTemplate#s3_configuration}
	S3Configuration *FisExperimentTemplateExperimentReportConfigurationOutputsS3Configuration `field:"optional" json:"s3Configuration" yaml:"s3Configuration"`
}

