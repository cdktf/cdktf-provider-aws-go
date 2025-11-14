// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package fmsresourceset


type FmsResourceSetResourceSet struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.21.0/docs/resources/fms_resource_set#name FmsResourceSet#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.21.0/docs/resources/fms_resource_set#description FmsResourceSet#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.21.0/docs/resources/fms_resource_set#resource_set_status FmsResourceSet#resource_set_status}.
	ResourceSetStatus *string `field:"optional" json:"resourceSetStatus" yaml:"resourceSetStatus"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.21.0/docs/resources/fms_resource_set#resource_type_list FmsResourceSet#resource_type_list}.
	ResourceTypeList *[]*string `field:"optional" json:"resourceTypeList" yaml:"resourceTypeList"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.21.0/docs/resources/fms_resource_set#update_token FmsResourceSet#update_token}.
	UpdateToken *string `field:"optional" json:"updateToken" yaml:"updateToken"`
}

