// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package configconfigrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ConfigConfigRuleConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/config_config_rule#name ConfigConfigRule#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/config_config_rule#source ConfigConfigRule#source}
	Source *ConfigConfigRuleSource `field:"required" json:"source" yaml:"source"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/config_config_rule#description ConfigConfigRule#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// evaluation_mode block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/config_config_rule#evaluation_mode ConfigConfigRule#evaluation_mode}
	EvaluationMode interface{} `field:"optional" json:"evaluationMode" yaml:"evaluationMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/config_config_rule#id ConfigConfigRule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/config_config_rule#input_parameters ConfigConfigRule#input_parameters}.
	InputParameters *string `field:"optional" json:"inputParameters" yaml:"inputParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/config_config_rule#maximum_execution_frequency ConfigConfigRule#maximum_execution_frequency}.
	MaximumExecutionFrequency *string `field:"optional" json:"maximumExecutionFrequency" yaml:"maximumExecutionFrequency"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/config_config_rule#region ConfigConfigRule#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// scope block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/config_config_rule#scope ConfigConfigRule#scope}
	Scope *ConfigConfigRuleScope `field:"optional" json:"scope" yaml:"scope"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/config_config_rule#tags ConfigConfigRule#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/config_config_rule#tags_all ConfigConfigRule#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
}

