// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package codepipeline


type CodepipelineStageBeforeEntryConditionRule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.91.0/docs/resources/codepipeline#name Codepipeline#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// rule_type_id block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.91.0/docs/resources/codepipeline#rule_type_id Codepipeline#rule_type_id}
	RuleTypeId *CodepipelineStageBeforeEntryConditionRuleRuleTypeId `field:"required" json:"ruleTypeId" yaml:"ruleTypeId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.91.0/docs/resources/codepipeline#commands Codepipeline#commands}.
	Commands *[]*string `field:"optional" json:"commands" yaml:"commands"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.91.0/docs/resources/codepipeline#configuration Codepipeline#configuration}.
	Configuration *map[string]*string `field:"optional" json:"configuration" yaml:"configuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.91.0/docs/resources/codepipeline#input_artifacts Codepipeline#input_artifacts}.
	InputArtifacts *[]*string `field:"optional" json:"inputArtifacts" yaml:"inputArtifacts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.91.0/docs/resources/codepipeline#region Codepipeline#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.91.0/docs/resources/codepipeline#role_arn Codepipeline#role_arn}.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.91.0/docs/resources/codepipeline#timeout_in_minutes Codepipeline#timeout_in_minutes}.
	TimeoutInMinutes *float64 `field:"optional" json:"timeoutInMinutes" yaml:"timeoutInMinutes"`
}

