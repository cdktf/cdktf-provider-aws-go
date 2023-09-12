// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package batchcomputeenvironment


type BatchComputeEnvironmentComputeResourcesEc2Configuration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.2/docs/resources/batch_compute_environment#image_id_override BatchComputeEnvironment#image_id_override}.
	ImageIdOverride *string `field:"optional" json:"imageIdOverride" yaml:"imageIdOverride"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.2/docs/resources/batch_compute_environment#image_type BatchComputeEnvironment#image_type}.
	ImageType *string `field:"optional" json:"imageType" yaml:"imageType"`
}

