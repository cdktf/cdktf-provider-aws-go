// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package fisexperimenttemplate


type FisExperimentTemplateTarget struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/fis_experiment_template#name FisExperimentTemplate#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/fis_experiment_template#resource_type FisExperimentTemplate#resource_type}.
	ResourceType *string `field:"required" json:"resourceType" yaml:"resourceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/fis_experiment_template#selection_mode FisExperimentTemplate#selection_mode}.
	SelectionMode *string `field:"required" json:"selectionMode" yaml:"selectionMode"`
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/fis_experiment_template#filter FisExperimentTemplate#filter}
	Filter interface{} `field:"optional" json:"filter" yaml:"filter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/fis_experiment_template#parameters FisExperimentTemplate#parameters}.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/fis_experiment_template#resource_arns FisExperimentTemplate#resource_arns}.
	ResourceArns *[]*string `field:"optional" json:"resourceArns" yaml:"resourceArns"`
	// resource_tag block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/fis_experiment_template#resource_tag FisExperimentTemplate#resource_tag}
	ResourceTag interface{} `field:"optional" json:"resourceTag" yaml:"resourceTag"`
}

