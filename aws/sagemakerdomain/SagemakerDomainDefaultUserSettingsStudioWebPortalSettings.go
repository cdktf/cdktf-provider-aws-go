// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakerdomain


type SagemakerDomainDefaultUserSettingsStudioWebPortalSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.88.0/docs/resources/sagemaker_domain#hidden_app_types SagemakerDomain#hidden_app_types}.
	HiddenAppTypes *[]*string `field:"optional" json:"hiddenAppTypes" yaml:"hiddenAppTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.88.0/docs/resources/sagemaker_domain#hidden_instance_types SagemakerDomain#hidden_instance_types}.
	HiddenInstanceTypes *[]*string `field:"optional" json:"hiddenInstanceTypes" yaml:"hiddenInstanceTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.88.0/docs/resources/sagemaker_domain#hidden_ml_tools SagemakerDomain#hidden_ml_tools}.
	HiddenMlTools *[]*string `field:"optional" json:"hiddenMlTools" yaml:"hiddenMlTools"`
}

