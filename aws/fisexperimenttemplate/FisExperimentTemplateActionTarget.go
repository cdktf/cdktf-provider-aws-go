// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package fisexperimenttemplate


type FisExperimentTemplateActionTarget struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/fis_experiment_template#key FisExperimentTemplate#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/fis_experiment_template#value FisExperimentTemplate#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

