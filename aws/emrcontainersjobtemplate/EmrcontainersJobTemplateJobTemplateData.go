// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package emrcontainersjobtemplate


type EmrcontainersJobTemplateJobTemplateData struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.70.0/docs/resources/emrcontainers_job_template#execution_role_arn EmrcontainersJobTemplate#execution_role_arn}.
	ExecutionRoleArn *string `field:"required" json:"executionRoleArn" yaml:"executionRoleArn"`
	// job_driver block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.70.0/docs/resources/emrcontainers_job_template#job_driver EmrcontainersJobTemplate#job_driver}
	JobDriver *EmrcontainersJobTemplateJobTemplateDataJobDriver `field:"required" json:"jobDriver" yaml:"jobDriver"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.70.0/docs/resources/emrcontainers_job_template#release_label EmrcontainersJobTemplate#release_label}.
	ReleaseLabel *string `field:"required" json:"releaseLabel" yaml:"releaseLabel"`
	// configuration_overrides block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.70.0/docs/resources/emrcontainers_job_template#configuration_overrides EmrcontainersJobTemplate#configuration_overrides}
	ConfigurationOverrides *EmrcontainersJobTemplateJobTemplateDataConfigurationOverrides `field:"optional" json:"configurationOverrides" yaml:"configurationOverrides"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.70.0/docs/resources/emrcontainers_job_template#job_tags EmrcontainersJobTemplate#job_tags}.
	JobTags *map[string]*string `field:"optional" json:"jobTags" yaml:"jobTags"`
}

