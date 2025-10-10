// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package servicecatalogproduct


type ServicecatalogProductProvisioningArtifactParameters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/servicecatalog_product#description ServicecatalogProduct#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/servicecatalog_product#disable_template_validation ServicecatalogProduct#disable_template_validation}.
	DisableTemplateValidation interface{} `field:"optional" json:"disableTemplateValidation" yaml:"disableTemplateValidation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/servicecatalog_product#name ServicecatalogProduct#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/servicecatalog_product#template_physical_id ServicecatalogProduct#template_physical_id}.
	TemplatePhysicalId *string `field:"optional" json:"templatePhysicalId" yaml:"templatePhysicalId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/servicecatalog_product#template_url ServicecatalogProduct#template_url}.
	TemplateUrl *string `field:"optional" json:"templateUrl" yaml:"templateUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/servicecatalog_product#type ServicecatalogProduct#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

