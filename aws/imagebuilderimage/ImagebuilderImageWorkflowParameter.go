// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package imagebuilderimage


type ImagebuilderImageWorkflowParameter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.13.0/docs/resources/imagebuilder_image#name ImagebuilderImage#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.13.0/docs/resources/imagebuilder_image#value ImagebuilderImage#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

