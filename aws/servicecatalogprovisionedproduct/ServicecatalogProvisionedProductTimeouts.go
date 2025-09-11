// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package servicecatalogprovisionedproduct


type ServicecatalogProvisionedProductTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.13.0/docs/resources/servicecatalog_provisioned_product#create ServicecatalogProvisionedProduct#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.13.0/docs/resources/servicecatalog_provisioned_product#delete ServicecatalogProvisionedProduct#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.13.0/docs/resources/servicecatalog_provisioned_product#read ServicecatalogProvisionedProduct#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.13.0/docs/resources/servicecatalog_provisioned_product#update ServicecatalogProvisionedProduct#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

