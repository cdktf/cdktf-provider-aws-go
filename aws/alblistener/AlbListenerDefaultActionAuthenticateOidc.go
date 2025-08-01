// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package alblistener


type AlbListenerDefaultActionAuthenticateOidc struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/alb_listener#authorization_endpoint AlbListener#authorization_endpoint}.
	AuthorizationEndpoint *string `field:"required" json:"authorizationEndpoint" yaml:"authorizationEndpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/alb_listener#client_id AlbListener#client_id}.
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/alb_listener#client_secret AlbListener#client_secret}.
	ClientSecret *string `field:"required" json:"clientSecret" yaml:"clientSecret"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/alb_listener#issuer AlbListener#issuer}.
	Issuer *string `field:"required" json:"issuer" yaml:"issuer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/alb_listener#token_endpoint AlbListener#token_endpoint}.
	TokenEndpoint *string `field:"required" json:"tokenEndpoint" yaml:"tokenEndpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/alb_listener#user_info_endpoint AlbListener#user_info_endpoint}.
	UserInfoEndpoint *string `field:"required" json:"userInfoEndpoint" yaml:"userInfoEndpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/alb_listener#authentication_request_extra_params AlbListener#authentication_request_extra_params}.
	AuthenticationRequestExtraParams *map[string]*string `field:"optional" json:"authenticationRequestExtraParams" yaml:"authenticationRequestExtraParams"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/alb_listener#on_unauthenticated_request AlbListener#on_unauthenticated_request}.
	OnUnauthenticatedRequest *string `field:"optional" json:"onUnauthenticatedRequest" yaml:"onUnauthenticatedRequest"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/alb_listener#scope AlbListener#scope}.
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/alb_listener#session_cookie_name AlbListener#session_cookie_name}.
	SessionCookieName *string `field:"optional" json:"sessionCookieName" yaml:"sessionCookieName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/alb_listener#session_timeout AlbListener#session_timeout}.
	SessionTimeout *float64 `field:"optional" json:"sessionTimeout" yaml:"sessionTimeout"`
}

