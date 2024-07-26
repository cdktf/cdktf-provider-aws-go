// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package batchjobdefinition


type BatchJobDefinitionEksPropertiesPodPropertiesVolumesSecret struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/batch_job_definition#secret_name BatchJobDefinition#secret_name}.
	SecretName *string `field:"required" json:"secretName" yaml:"secretName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/batch_job_definition#optional BatchJobDefinition#optional}.
	Optional interface{} `field:"optional" json:"optional" yaml:"optional"`
}

