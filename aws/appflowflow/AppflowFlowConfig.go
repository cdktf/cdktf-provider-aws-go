// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appflowflow

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppflowFlowConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// destination_flow_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/appflow_flow#destination_flow_config AppflowFlow#destination_flow_config}
	DestinationFlowConfig interface{} `field:"required" json:"destinationFlowConfig" yaml:"destinationFlowConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/appflow_flow#name AppflowFlow#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// source_flow_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/appflow_flow#source_flow_config AppflowFlow#source_flow_config}
	SourceFlowConfig *AppflowFlowSourceFlowConfig `field:"required" json:"sourceFlowConfig" yaml:"sourceFlowConfig"`
	// task block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/appflow_flow#task AppflowFlow#task}
	Task interface{} `field:"required" json:"task" yaml:"task"`
	// trigger_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/appflow_flow#trigger_config AppflowFlow#trigger_config}
	TriggerConfig *AppflowFlowTriggerConfig `field:"required" json:"triggerConfig" yaml:"triggerConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/appflow_flow#description AppflowFlow#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/appflow_flow#id AppflowFlow#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/appflow_flow#kms_arn AppflowFlow#kms_arn}.
	KmsArn *string `field:"optional" json:"kmsArn" yaml:"kmsArn"`
	// metadata_catalog_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/appflow_flow#metadata_catalog_config AppflowFlow#metadata_catalog_config}
	MetadataCatalogConfig *AppflowFlowMetadataCatalogConfig `field:"optional" json:"metadataCatalogConfig" yaml:"metadataCatalogConfig"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/appflow_flow#region AppflowFlow#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/appflow_flow#tags AppflowFlow#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/appflow_flow#tags_all AppflowFlow#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
}

