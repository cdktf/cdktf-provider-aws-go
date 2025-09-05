// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package imagebuilderlifecyclepolicy


type ImagebuilderLifecyclePolicyResourceSelectionRecipe struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.12.0/docs/resources/imagebuilder_lifecycle_policy#name ImagebuilderLifecyclePolicy#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.12.0/docs/resources/imagebuilder_lifecycle_policy#semantic_version ImagebuilderLifecyclePolicy#semantic_version}.
	SemanticVersion *string `field:"required" json:"semanticVersion" yaml:"semanticVersion"`
}

