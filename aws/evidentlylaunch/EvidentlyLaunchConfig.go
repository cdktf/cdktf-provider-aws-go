// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package evidentlylaunch

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EvidentlyLaunchConfig struct {
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
	// groups block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/evidently_launch#groups EvidentlyLaunch#groups}
	Groups interface{} `field:"required" json:"groups" yaml:"groups"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/evidently_launch#name EvidentlyLaunch#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/evidently_launch#project EvidentlyLaunch#project}.
	Project *string `field:"required" json:"project" yaml:"project"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/evidently_launch#description EvidentlyLaunch#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/evidently_launch#id EvidentlyLaunch#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// metric_monitors block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/evidently_launch#metric_monitors EvidentlyLaunch#metric_monitors}
	MetricMonitors interface{} `field:"optional" json:"metricMonitors" yaml:"metricMonitors"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/evidently_launch#randomization_salt EvidentlyLaunch#randomization_salt}.
	RandomizationSalt *string `field:"optional" json:"randomizationSalt" yaml:"randomizationSalt"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/evidently_launch#region EvidentlyLaunch#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// scheduled_splits_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/evidently_launch#scheduled_splits_config EvidentlyLaunch#scheduled_splits_config}
	ScheduledSplitsConfig *EvidentlyLaunchScheduledSplitsConfig `field:"optional" json:"scheduledSplitsConfig" yaml:"scheduledSplitsConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/evidently_launch#tags EvidentlyLaunch#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/evidently_launch#tags_all EvidentlyLaunch#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/evidently_launch#timeouts EvidentlyLaunch#timeouts}
	Timeouts *EvidentlyLaunchTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

