// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package fisexperimenttemplate


type FisExperimentTemplateExperimentReportConfiguration struct {
	// data_sources block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.11.0/docs/resources/fis_experiment_template#data_sources FisExperimentTemplate#data_sources}
	DataSources *FisExperimentTemplateExperimentReportConfigurationDataSources `field:"optional" json:"dataSources" yaml:"dataSources"`
	// outputs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.11.0/docs/resources/fis_experiment_template#outputs FisExperimentTemplate#outputs}
	Outputs *FisExperimentTemplateExperimentReportConfigurationOutputs `field:"optional" json:"outputs" yaml:"outputs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.11.0/docs/resources/fis_experiment_template#post_experiment_duration FisExperimentTemplate#post_experiment_duration}.
	PostExperimentDuration *string `field:"optional" json:"postExperimentDuration" yaml:"postExperimentDuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.11.0/docs/resources/fis_experiment_template#pre_experiment_duration FisExperimentTemplate#pre_experiment_duration}.
	PreExperimentDuration *string `field:"optional" json:"preExperimentDuration" yaml:"preExperimentDuration"`
}

