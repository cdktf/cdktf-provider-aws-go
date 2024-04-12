// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package alblistener


type AlbListenerDefaultActionAuthenticateCognito struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/alb_listener#user_pool_arn AlbListener#user_pool_arn}.
	UserPoolArn *string `field:"required" json:"userPoolArn" yaml:"userPoolArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/alb_listener#user_pool_client_id AlbListener#user_pool_client_id}.
	UserPoolClientId *string `field:"required" json:"userPoolClientId" yaml:"userPoolClientId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/alb_listener#user_pool_domain AlbListener#user_pool_domain}.
	UserPoolDomain *string `field:"required" json:"userPoolDomain" yaml:"userPoolDomain"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/alb_listener#authentication_request_extra_params AlbListener#authentication_request_extra_params}.
	AuthenticationRequestExtraParams *map[string]*string `field:"optional" json:"authenticationRequestExtraParams" yaml:"authenticationRequestExtraParams"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/alb_listener#on_unauthenticated_request AlbListener#on_unauthenticated_request}.
	OnUnauthenticatedRequest *string `field:"optional" json:"onUnauthenticatedRequest" yaml:"onUnauthenticatedRequest"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/alb_listener#scope AlbListener#scope}.
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/alb_listener#session_cookie_name AlbListener#session_cookie_name}.
	SessionCookieName *string `field:"optional" json:"sessionCookieName" yaml:"sessionCookieName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/alb_listener#session_timeout AlbListener#session_timeout}.
	SessionTimeout *float64 `field:"optional" json:"sessionTimeout" yaml:"sessionTimeout"`
}

