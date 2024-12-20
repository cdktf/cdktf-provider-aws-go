// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockagentknowledgebase

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BedrockagentKnowledgeBaseConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/resources/bedrockagent_knowledge_base#name BedrockagentKnowledgeBase#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/resources/bedrockagent_knowledge_base#role_arn BedrockagentKnowledgeBase#role_arn}.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/resources/bedrockagent_knowledge_base#description BedrockagentKnowledgeBase#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// knowledge_base_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/resources/bedrockagent_knowledge_base#knowledge_base_configuration BedrockagentKnowledgeBase#knowledge_base_configuration}
	KnowledgeBaseConfiguration interface{} `field:"optional" json:"knowledgeBaseConfiguration" yaml:"knowledgeBaseConfiguration"`
	// storage_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/resources/bedrockagent_knowledge_base#storage_configuration BedrockagentKnowledgeBase#storage_configuration}
	StorageConfiguration interface{} `field:"optional" json:"storageConfiguration" yaml:"storageConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/resources/bedrockagent_knowledge_base#tags BedrockagentKnowledgeBase#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/resources/bedrockagent_knowledge_base#timeouts BedrockagentKnowledgeBase#timeouts}
	Timeouts *BedrockagentKnowledgeBaseTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

