// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gluejob


type GlueJobCommand struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/glue_job#script_location GlueJob#script_location}.
	ScriptLocation *string `field:"required" json:"scriptLocation" yaml:"scriptLocation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/glue_job#name GlueJob#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/glue_job#python_version GlueJob#python_version}.
	PythonVersion *string `field:"optional" json:"pythonVersion" yaml:"pythonVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/glue_job#runtime GlueJob#runtime}.
	Runtime *string `field:"optional" json:"runtime" yaml:"runtime"`
}

