// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cognitomanageduserpoolclient


type CognitoManagedUserPoolClientAnalyticsConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/cognito_managed_user_pool_client#application_arn CognitoManagedUserPoolClient#application_arn}.
	ApplicationArn *string `field:"optional" json:"applicationArn" yaml:"applicationArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/cognito_managed_user_pool_client#application_id CognitoManagedUserPoolClient#application_id}.
	ApplicationId *string `field:"optional" json:"applicationId" yaml:"applicationId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/cognito_managed_user_pool_client#external_id CognitoManagedUserPoolClient#external_id}.
	ExternalId *string `field:"optional" json:"externalId" yaml:"externalId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/cognito_managed_user_pool_client#role_arn CognitoManagedUserPoolClient#role_arn}.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/cognito_managed_user_pool_client#user_data_shared CognitoManagedUserPoolClient#user_data_shared}.
	UserDataShared interface{} `field:"optional" json:"userDataShared" yaml:"userDataShared"`
}

