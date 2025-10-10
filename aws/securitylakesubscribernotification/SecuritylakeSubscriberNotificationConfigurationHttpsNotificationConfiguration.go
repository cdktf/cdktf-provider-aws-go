// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package securitylakesubscribernotification


type SecuritylakeSubscriberNotificationConfigurationHttpsNotificationConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/securitylake_subscriber_notification#endpoint SecuritylakeSubscriberNotification#endpoint}.
	Endpoint *string `field:"required" json:"endpoint" yaml:"endpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/securitylake_subscriber_notification#target_role_arn SecuritylakeSubscriberNotification#target_role_arn}.
	TargetRoleArn *string `field:"required" json:"targetRoleArn" yaml:"targetRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/securitylake_subscriber_notification#authorization_api_key_name SecuritylakeSubscriberNotification#authorization_api_key_name}.
	AuthorizationApiKeyName *string `field:"optional" json:"authorizationApiKeyName" yaml:"authorizationApiKeyName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/securitylake_subscriber_notification#authorization_api_key_value SecuritylakeSubscriberNotification#authorization_api_key_value}.
	AuthorizationApiKeyValue *string `field:"optional" json:"authorizationApiKeyValue" yaml:"authorizationApiKeyValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/securitylake_subscriber_notification#http_method SecuritylakeSubscriberNotification#http_method}.
	HttpMethod *string `field:"optional" json:"httpMethod" yaml:"httpMethod"`
}

