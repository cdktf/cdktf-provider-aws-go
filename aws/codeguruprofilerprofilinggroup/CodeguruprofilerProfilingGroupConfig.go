// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package codeguruprofilerprofilinggroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CodeguruprofilerProfilingGroupConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/resources/codeguruprofiler_profiling_group#name CodeguruprofilerProfilingGroup#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// agent_orchestration_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/resources/codeguruprofiler_profiling_group#agent_orchestration_config CodeguruprofilerProfilingGroup#agent_orchestration_config}
	AgentOrchestrationConfig interface{} `field:"optional" json:"agentOrchestrationConfig" yaml:"agentOrchestrationConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/resources/codeguruprofiler_profiling_group#compute_platform CodeguruprofilerProfilingGroup#compute_platform}.
	ComputePlatform *string `field:"optional" json:"computePlatform" yaml:"computePlatform"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/resources/codeguruprofiler_profiling_group#tags CodeguruprofilerProfilingGroup#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

