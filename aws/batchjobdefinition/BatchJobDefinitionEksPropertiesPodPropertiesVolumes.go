// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package batchjobdefinition


type BatchJobDefinitionEksPropertiesPodPropertiesVolumes struct {
	// empty_dir block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.52.0/docs/resources/batch_job_definition#empty_dir BatchJobDefinition#empty_dir}
	EmptyDir *BatchJobDefinitionEksPropertiesPodPropertiesVolumesEmptyDir `field:"optional" json:"emptyDir" yaml:"emptyDir"`
	// host_path block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.52.0/docs/resources/batch_job_definition#host_path BatchJobDefinition#host_path}
	HostPath *BatchJobDefinitionEksPropertiesPodPropertiesVolumesHostPath `field:"optional" json:"hostPath" yaml:"hostPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.52.0/docs/resources/batch_job_definition#name BatchJobDefinition#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// secret block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.52.0/docs/resources/batch_job_definition#secret BatchJobDefinition#secret}
	Secret *BatchJobDefinitionEksPropertiesPodPropertiesVolumesSecret `field:"optional" json:"secret" yaml:"secret"`
}

