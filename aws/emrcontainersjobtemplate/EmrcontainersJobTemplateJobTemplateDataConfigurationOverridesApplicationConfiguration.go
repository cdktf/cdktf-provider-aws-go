// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package emrcontainersjobtemplate


type EmrcontainersJobTemplateJobTemplateDataConfigurationOverridesApplicationConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.1/docs/resources/emrcontainers_job_template#classification EmrcontainersJobTemplate#classification}.
	Classification *string `field:"required" json:"classification" yaml:"classification"`
	// configurations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.1/docs/resources/emrcontainers_job_template#configurations EmrcontainersJobTemplate#configurations}
	Configurations interface{} `field:"optional" json:"configurations" yaml:"configurations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.1/docs/resources/emrcontainers_job_template#properties EmrcontainersJobTemplate#properties}.
	Properties *map[string]*string `field:"optional" json:"properties" yaml:"properties"`
}

