// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package imagebuilderlifecyclepolicy


type ImagebuilderLifecyclePolicyResourceSelection struct {
	// recipe block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/imagebuilder_lifecycle_policy#recipe ImagebuilderLifecyclePolicy#recipe}
	Recipe interface{} `field:"optional" json:"recipe" yaml:"recipe"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/imagebuilder_lifecycle_policy#tag_map ImagebuilderLifecyclePolicy#tag_map}.
	TagMap *map[string]*string `field:"optional" json:"tagMap" yaml:"tagMap"`
}

