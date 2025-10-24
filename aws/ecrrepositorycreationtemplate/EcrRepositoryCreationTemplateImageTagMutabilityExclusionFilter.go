// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ecrrepositorycreationtemplate


type EcrRepositoryCreationTemplateImageTagMutabilityExclusionFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.18.0/docs/resources/ecr_repository_creation_template#filter EcrRepositoryCreationTemplate#filter}.
	Filter *string `field:"required" json:"filter" yaml:"filter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.18.0/docs/resources/ecr_repository_creation_template#filter_type EcrRepositoryCreationTemplate#filter_type}.
	FilterType *string `field:"required" json:"filterType" yaml:"filterType"`
}

