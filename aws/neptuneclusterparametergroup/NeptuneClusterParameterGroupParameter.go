// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package neptuneclusterparametergroup


type NeptuneClusterParameterGroupParameter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.79.0/docs/resources/neptune_cluster_parameter_group#name NeptuneClusterParameterGroup#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.79.0/docs/resources/neptune_cluster_parameter_group#value NeptuneClusterParameterGroup#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.79.0/docs/resources/neptune_cluster_parameter_group#apply_method NeptuneClusterParameterGroup#apply_method}.
	ApplyMethod *string `field:"optional" json:"applyMethod" yaml:"applyMethod"`
}

