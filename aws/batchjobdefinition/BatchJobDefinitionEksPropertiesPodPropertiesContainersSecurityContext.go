// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package batchjobdefinition


type BatchJobDefinitionEksPropertiesPodPropertiesContainersSecurityContext struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/batch_job_definition#privileged BatchJobDefinition#privileged}.
	Privileged interface{} `field:"optional" json:"privileged" yaml:"privileged"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/batch_job_definition#read_only_root_file_system BatchJobDefinition#read_only_root_file_system}.
	ReadOnlyRootFileSystem interface{} `field:"optional" json:"readOnlyRootFileSystem" yaml:"readOnlyRootFileSystem"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/batch_job_definition#run_as_group BatchJobDefinition#run_as_group}.
	RunAsGroup *float64 `field:"optional" json:"runAsGroup" yaml:"runAsGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/batch_job_definition#run_as_non_root BatchJobDefinition#run_as_non_root}.
	RunAsNonRoot interface{} `field:"optional" json:"runAsNonRoot" yaml:"runAsNonRoot"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/batch_job_definition#run_as_user BatchJobDefinition#run_as_user}.
	RunAsUser *float64 `field:"optional" json:"runAsUser" yaml:"runAsUser"`
}

