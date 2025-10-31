// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datazoneenvironment


type DatazoneEnvironmentUserParameters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/datazone_environment#name DatazoneEnvironment#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/datazone_environment#value DatazoneEnvironment#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

