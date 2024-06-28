// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package codepipelinecustomactiontype


type CodepipelineCustomActionTypeSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.56.0/docs/resources/codepipeline_custom_action_type#entity_url_template CodepipelineCustomActionType#entity_url_template}.
	EntityUrlTemplate *string `field:"optional" json:"entityUrlTemplate" yaml:"entityUrlTemplate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.56.0/docs/resources/codepipeline_custom_action_type#execution_url_template CodepipelineCustomActionType#execution_url_template}.
	ExecutionUrlTemplate *string `field:"optional" json:"executionUrlTemplate" yaml:"executionUrlTemplate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.56.0/docs/resources/codepipeline_custom_action_type#revision_url_template CodepipelineCustomActionType#revision_url_template}.
	RevisionUrlTemplate *string `field:"optional" json:"revisionUrlTemplate" yaml:"revisionUrlTemplate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.56.0/docs/resources/codepipeline_custom_action_type#third_party_configuration_url CodepipelineCustomActionType#third_party_configuration_url}.
	ThirdPartyConfigurationUrl *string `field:"optional" json:"thirdPartyConfigurationUrl" yaml:"thirdPartyConfigurationUrl"`
}

