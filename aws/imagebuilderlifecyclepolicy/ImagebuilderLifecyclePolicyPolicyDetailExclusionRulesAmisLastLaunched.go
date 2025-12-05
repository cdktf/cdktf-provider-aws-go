// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package imagebuilderlifecyclepolicy


type ImagebuilderLifecyclePolicyPolicyDetailExclusionRulesAmisLastLaunched struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/imagebuilder_lifecycle_policy#unit ImagebuilderLifecyclePolicy#unit}.
	Unit *string `field:"required" json:"unit" yaml:"unit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/imagebuilder_lifecycle_policy#value ImagebuilderLifecyclePolicy#value}.
	Value *float64 `field:"required" json:"value" yaml:"value"`
}

