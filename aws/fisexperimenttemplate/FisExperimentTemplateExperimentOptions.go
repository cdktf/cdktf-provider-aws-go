// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package fisexperimenttemplate


type FisExperimentTemplateExperimentOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/fis_experiment_template#account_targeting FisExperimentTemplate#account_targeting}.
	AccountTargeting *string `field:"optional" json:"accountTargeting" yaml:"accountTargeting"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/fis_experiment_template#empty_target_resolution_mode FisExperimentTemplate#empty_target_resolution_mode}.
	EmptyTargetResolutionMode *string `field:"optional" json:"emptyTargetResolutionMode" yaml:"emptyTargetResolutionMode"`
}

