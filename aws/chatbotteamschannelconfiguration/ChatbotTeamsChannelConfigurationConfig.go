// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package chatbotteamschannelconfiguration

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ChatbotTeamsChannelConfigurationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/chatbot_teams_channel_configuration#channel_id ChatbotTeamsChannelConfiguration#channel_id}.
	ChannelId *string `field:"required" json:"channelId" yaml:"channelId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/chatbot_teams_channel_configuration#configuration_name ChatbotTeamsChannelConfiguration#configuration_name}.
	ConfigurationName *string `field:"required" json:"configurationName" yaml:"configurationName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/chatbot_teams_channel_configuration#iam_role_arn ChatbotTeamsChannelConfiguration#iam_role_arn}.
	IamRoleArn *string `field:"required" json:"iamRoleArn" yaml:"iamRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/chatbot_teams_channel_configuration#team_id ChatbotTeamsChannelConfiguration#team_id}.
	TeamId *string `field:"required" json:"teamId" yaml:"teamId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/chatbot_teams_channel_configuration#tenant_id ChatbotTeamsChannelConfiguration#tenant_id}.
	TenantId *string `field:"required" json:"tenantId" yaml:"tenantId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/chatbot_teams_channel_configuration#channel_name ChatbotTeamsChannelConfiguration#channel_name}.
	ChannelName *string `field:"optional" json:"channelName" yaml:"channelName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/chatbot_teams_channel_configuration#guardrail_policy_arns ChatbotTeamsChannelConfiguration#guardrail_policy_arns}.
	GuardrailPolicyArns *[]*string `field:"optional" json:"guardrailPolicyArns" yaml:"guardrailPolicyArns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/chatbot_teams_channel_configuration#logging_level ChatbotTeamsChannelConfiguration#logging_level}.
	LoggingLevel *string `field:"optional" json:"loggingLevel" yaml:"loggingLevel"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/chatbot_teams_channel_configuration#region ChatbotTeamsChannelConfiguration#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/chatbot_teams_channel_configuration#sns_topic_arns ChatbotTeamsChannelConfiguration#sns_topic_arns}.
	SnsTopicArns *[]*string `field:"optional" json:"snsTopicArns" yaml:"snsTopicArns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/chatbot_teams_channel_configuration#tags ChatbotTeamsChannelConfiguration#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/chatbot_teams_channel_configuration#team_name ChatbotTeamsChannelConfiguration#team_name}.
	TeamName *string `field:"optional" json:"teamName" yaml:"teamName"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/chatbot_teams_channel_configuration#timeouts ChatbotTeamsChannelConfiguration#timeouts}
	Timeouts *ChatbotTeamsChannelConfigurationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/chatbot_teams_channel_configuration#user_authorization_required ChatbotTeamsChannelConfiguration#user_authorization_required}.
	UserAuthorizationRequired interface{} `field:"optional" json:"userAuthorizationRequired" yaml:"userAuthorizationRequired"`
}

