// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package imagebuilderlifecyclepolicy


type ImagebuilderLifecyclePolicyPolicyDetailActionIncludeResources struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/imagebuilder_lifecycle_policy#amis ImagebuilderLifecyclePolicy#amis}.
	Amis interface{} `field:"optional" json:"amis" yaml:"amis"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/imagebuilder_lifecycle_policy#containers ImagebuilderLifecyclePolicy#containers}.
	Containers interface{} `field:"optional" json:"containers" yaml:"containers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/imagebuilder_lifecycle_policy#snapshots ImagebuilderLifecyclePolicy#snapshots}.
	Snapshots interface{} `field:"optional" json:"snapshots" yaml:"snapshots"`
}

