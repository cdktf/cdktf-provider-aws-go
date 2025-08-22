// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockagentflow


type BedrockagentFlowDefinitionNodeConfiguration struct {
	// agent block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/bedrockagent_flow#agent BedrockagentFlow#agent}
	Agent interface{} `field:"optional" json:"agent" yaml:"agent"`
	// collector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/bedrockagent_flow#collector BedrockagentFlow#collector}
	Collector interface{} `field:"optional" json:"collector" yaml:"collector"`
	// condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/bedrockagent_flow#condition BedrockagentFlow#condition}
	Condition interface{} `field:"optional" json:"condition" yaml:"condition"`
	// inline_code block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/bedrockagent_flow#inline_code BedrockagentFlow#inline_code}
	InlineCode interface{} `field:"optional" json:"inlineCode" yaml:"inlineCode"`
	// input block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/bedrockagent_flow#input BedrockagentFlow#input}
	Input interface{} `field:"optional" json:"input" yaml:"input"`
	// iterator block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/bedrockagent_flow#iterator BedrockagentFlow#iterator}
	Iterator interface{} `field:"optional" json:"iterator" yaml:"iterator"`
	// knowledge_base block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/bedrockagent_flow#knowledge_base BedrockagentFlow#knowledge_base}
	KnowledgeBase interface{} `field:"optional" json:"knowledgeBase" yaml:"knowledgeBase"`
	// lambda_function block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/bedrockagent_flow#lambda_function BedrockagentFlow#lambda_function}
	LambdaFunction interface{} `field:"optional" json:"lambdaFunction" yaml:"lambdaFunction"`
	// lex block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/bedrockagent_flow#lex BedrockagentFlow#lex}
	Lex interface{} `field:"optional" json:"lex" yaml:"lex"`
	// output block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/bedrockagent_flow#output BedrockagentFlow#output}
	Output interface{} `field:"optional" json:"output" yaml:"output"`
	// prompt block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/bedrockagent_flow#prompt BedrockagentFlow#prompt}
	Prompt interface{} `field:"optional" json:"prompt" yaml:"prompt"`
	// retrieval block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/bedrockagent_flow#retrieval BedrockagentFlow#retrieval}
	Retrieval interface{} `field:"optional" json:"retrieval" yaml:"retrieval"`
	// storage block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/bedrockagent_flow#storage BedrockagentFlow#storage}
	Storage interface{} `field:"optional" json:"storage" yaml:"storage"`
}

