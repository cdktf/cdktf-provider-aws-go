// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package imagebuilderlifecyclepolicy


type ImagebuilderLifecyclePolicyPolicyDetailExclusionRules struct {
	// amis block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/imagebuilder_lifecycle_policy#amis ImagebuilderLifecyclePolicy#amis}
	Amis interface{} `field:"optional" json:"amis" yaml:"amis"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/imagebuilder_lifecycle_policy#tag_map ImagebuilderLifecyclePolicy#tag_map}.
	TagMap *map[string]*string `field:"optional" json:"tagMap" yaml:"tagMap"`
}

