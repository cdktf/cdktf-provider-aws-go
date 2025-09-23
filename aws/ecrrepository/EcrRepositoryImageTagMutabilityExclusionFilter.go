// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ecrrepository


type EcrRepositoryImageTagMutabilityExclusionFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/ecr_repository#filter EcrRepository#filter}.
	Filter *string `field:"required" json:"filter" yaml:"filter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/ecr_repository#filter_type EcrRepository#filter_type}.
	FilterType *string `field:"required" json:"filterType" yaml:"filterType"`
}

