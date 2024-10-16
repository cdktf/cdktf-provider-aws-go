// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package daxparametergroup


type DaxParameterGroupParameters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.72.0/docs/resources/dax_parameter_group#name DaxParameterGroup#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.72.0/docs/resources/dax_parameter_group#value DaxParameterGroup#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

