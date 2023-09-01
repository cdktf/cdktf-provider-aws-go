// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakerproject


type SagemakerProjectServiceCatalogProvisioningDetailsProvisioningParameter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/sagemaker_project#key SagemakerProject#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/sagemaker_project#value SagemakerProject#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

