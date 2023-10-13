// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package servicecatalogprovisionedproduct


type ServicecatalogProvisionedProductProvisioningParameters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.21.0/docs/resources/servicecatalog_provisioned_product#key ServicecatalogProvisionedProduct#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.21.0/docs/resources/servicecatalog_provisioned_product#use_previous_value ServicecatalogProvisionedProduct#use_previous_value}.
	UsePreviousValue interface{} `field:"optional" json:"usePreviousValue" yaml:"usePreviousValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.21.0/docs/resources/servicecatalog_provisioned_product#value ServicecatalogProvisionedProduct#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

