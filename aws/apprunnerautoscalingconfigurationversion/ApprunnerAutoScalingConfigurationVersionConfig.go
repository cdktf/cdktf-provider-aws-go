// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apprunnerautoscalingconfigurationversion

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApprunnerAutoScalingConfigurationVersionConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.30.0/docs/resources/apprunner_auto_scaling_configuration_version#auto_scaling_configuration_name ApprunnerAutoScalingConfigurationVersion#auto_scaling_configuration_name}.
	AutoScalingConfigurationName *string `field:"required" json:"autoScalingConfigurationName" yaml:"autoScalingConfigurationName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.30.0/docs/resources/apprunner_auto_scaling_configuration_version#id ApprunnerAutoScalingConfigurationVersion#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.30.0/docs/resources/apprunner_auto_scaling_configuration_version#max_concurrency ApprunnerAutoScalingConfigurationVersion#max_concurrency}.
	MaxConcurrency *float64 `field:"optional" json:"maxConcurrency" yaml:"maxConcurrency"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.30.0/docs/resources/apprunner_auto_scaling_configuration_version#max_size ApprunnerAutoScalingConfigurationVersion#max_size}.
	MaxSize *float64 `field:"optional" json:"maxSize" yaml:"maxSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.30.0/docs/resources/apprunner_auto_scaling_configuration_version#min_size ApprunnerAutoScalingConfigurationVersion#min_size}.
	MinSize *float64 `field:"optional" json:"minSize" yaml:"minSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.30.0/docs/resources/apprunner_auto_scaling_configuration_version#tags ApprunnerAutoScalingConfigurationVersion#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.30.0/docs/resources/apprunner_auto_scaling_configuration_version#tags_all ApprunnerAutoScalingConfigurationVersion#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
}

