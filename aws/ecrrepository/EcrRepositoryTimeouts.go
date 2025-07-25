// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ecrrepository


type EcrRepositoryTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/ecr_repository#delete EcrRepository#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

