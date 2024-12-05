// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package servicecatalogserviceaction


type ServicecatalogServiceActionDefinition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.80.0/docs/resources/servicecatalog_service_action#name ServicecatalogServiceAction#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.80.0/docs/resources/servicecatalog_service_action#version ServicecatalogServiceAction#version}.
	Version *string `field:"required" json:"version" yaml:"version"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.80.0/docs/resources/servicecatalog_service_action#assume_role ServicecatalogServiceAction#assume_role}.
	AssumeRole *string `field:"optional" json:"assumeRole" yaml:"assumeRole"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.80.0/docs/resources/servicecatalog_service_action#parameters ServicecatalogServiceAction#parameters}.
	Parameters *string `field:"optional" json:"parameters" yaml:"parameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.80.0/docs/resources/servicecatalog_service_action#type ServicecatalogServiceAction#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

