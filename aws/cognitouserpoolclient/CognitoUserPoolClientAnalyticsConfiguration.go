// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cognitouserpoolclient


type CognitoUserPoolClientAnalyticsConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.1/docs/resources/cognito_user_pool_client#application_arn CognitoUserPoolClient#application_arn}.
	ApplicationArn *string `field:"optional" json:"applicationArn" yaml:"applicationArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.1/docs/resources/cognito_user_pool_client#application_id CognitoUserPoolClient#application_id}.
	ApplicationId *string `field:"optional" json:"applicationId" yaml:"applicationId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.1/docs/resources/cognito_user_pool_client#external_id CognitoUserPoolClient#external_id}.
	ExternalId *string `field:"optional" json:"externalId" yaml:"externalId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.1/docs/resources/cognito_user_pool_client#role_arn CognitoUserPoolClient#role_arn}.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.1/docs/resources/cognito_user_pool_client#user_data_shared CognitoUserPoolClient#user_data_shared}.
	UserDataShared interface{} `field:"optional" json:"userDataShared" yaml:"userDataShared"`
}

