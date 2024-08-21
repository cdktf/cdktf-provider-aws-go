// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dbparametergroup


type DbParameterGroupParameter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.1/docs/resources/db_parameter_group#name DbParameterGroup#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.1/docs/resources/db_parameter_group#value DbParameterGroup#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.1/docs/resources/db_parameter_group#apply_method DbParameterGroup#apply_method}.
	ApplyMethod *string `field:"optional" json:"applyMethod" yaml:"applyMethod"`
}

