// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package imagebuilderlifecyclepolicy


type ImagebuilderLifecyclePolicyPolicyDetailAction struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/imagebuilder_lifecycle_policy#type ImagebuilderLifecyclePolicy#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// include_resources block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/imagebuilder_lifecycle_policy#include_resources ImagebuilderLifecyclePolicy#include_resources}
	IncludeResources interface{} `field:"optional" json:"includeResources" yaml:"includeResources"`
}

