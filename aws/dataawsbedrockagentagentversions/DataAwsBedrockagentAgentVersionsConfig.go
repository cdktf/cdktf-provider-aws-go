// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsbedrockagentagentversions

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAwsBedrockagentAgentVersionsConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/data-sources/bedrockagent_agent_versions#agent_id DataAwsBedrockagentAgentVersions#agent_id}.
	AgentId *string `field:"required" json:"agentId" yaml:"agentId"`
	// agent_version_summaries block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/data-sources/bedrockagent_agent_versions#agent_version_summaries DataAwsBedrockagentAgentVersions#agent_version_summaries}
	AgentVersionSummaries interface{} `field:"optional" json:"agentVersionSummaries" yaml:"agentVersionSummaries"`
}

