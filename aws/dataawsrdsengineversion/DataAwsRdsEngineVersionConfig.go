// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsrdsengineversion

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAwsRdsEngineVersionConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.12.0/docs/data-sources/rds_engine_version#engine DataAwsRdsEngineVersion#engine}.
	Engine *string `field:"required" json:"engine" yaml:"engine"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.12.0/docs/data-sources/rds_engine_version#default_only DataAwsRdsEngineVersion#default_only}.
	DefaultOnly interface{} `field:"optional" json:"defaultOnly" yaml:"defaultOnly"`
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.12.0/docs/data-sources/rds_engine_version#filter DataAwsRdsEngineVersion#filter}
	Filter interface{} `field:"optional" json:"filter" yaml:"filter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.12.0/docs/data-sources/rds_engine_version#has_major_target DataAwsRdsEngineVersion#has_major_target}.
	HasMajorTarget interface{} `field:"optional" json:"hasMajorTarget" yaml:"hasMajorTarget"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.12.0/docs/data-sources/rds_engine_version#has_minor_target DataAwsRdsEngineVersion#has_minor_target}.
	HasMinorTarget interface{} `field:"optional" json:"hasMinorTarget" yaml:"hasMinorTarget"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.12.0/docs/data-sources/rds_engine_version#id DataAwsRdsEngineVersion#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.12.0/docs/data-sources/rds_engine_version#include_all DataAwsRdsEngineVersion#include_all}.
	IncludeAll interface{} `field:"optional" json:"includeAll" yaml:"includeAll"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.12.0/docs/data-sources/rds_engine_version#latest DataAwsRdsEngineVersion#latest}.
	Latest interface{} `field:"optional" json:"latest" yaml:"latest"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.12.0/docs/data-sources/rds_engine_version#parameter_group_family DataAwsRdsEngineVersion#parameter_group_family}.
	ParameterGroupFamily *string `field:"optional" json:"parameterGroupFamily" yaml:"parameterGroupFamily"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.12.0/docs/data-sources/rds_engine_version#preferred_major_targets DataAwsRdsEngineVersion#preferred_major_targets}.
	PreferredMajorTargets *[]*string `field:"optional" json:"preferredMajorTargets" yaml:"preferredMajorTargets"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.12.0/docs/data-sources/rds_engine_version#preferred_upgrade_targets DataAwsRdsEngineVersion#preferred_upgrade_targets}.
	PreferredUpgradeTargets *[]*string `field:"optional" json:"preferredUpgradeTargets" yaml:"preferredUpgradeTargets"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.12.0/docs/data-sources/rds_engine_version#preferred_versions DataAwsRdsEngineVersion#preferred_versions}.
	PreferredVersions *[]*string `field:"optional" json:"preferredVersions" yaml:"preferredVersions"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.12.0/docs/data-sources/rds_engine_version#region DataAwsRdsEngineVersion#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.12.0/docs/data-sources/rds_engine_version#version DataAwsRdsEngineVersion#version}.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

