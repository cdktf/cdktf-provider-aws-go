// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakerworkteam


type SagemakerWorkteamMemberDefinitionCognitoMemberDefinition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.0/docs/resources/sagemaker_workteam#client_id SagemakerWorkteam#client_id}.
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.0/docs/resources/sagemaker_workteam#user_group SagemakerWorkteam#user_group}.
	UserGroup *string `field:"required" json:"userGroup" yaml:"userGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.0/docs/resources/sagemaker_workteam#user_pool SagemakerWorkteam#user_pool}.
	UserPool *string `field:"required" json:"userPool" yaml:"userPool"`
}

