// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package batchjobdefinition


type BatchJobDefinitionEksPropertiesPodPropertiesVolumesEmptyDir struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.68.0/docs/resources/batch_job_definition#size_limit BatchJobDefinition#size_limit}.
	SizeLimit *string `field:"required" json:"sizeLimit" yaml:"sizeLimit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.68.0/docs/resources/batch_job_definition#medium BatchJobDefinition#medium}.
	Medium *string `field:"optional" json:"medium" yaml:"medium"`
}

