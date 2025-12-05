// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lblistener


type LbListenerDefaultActionAuthenticateOidc struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/lb_listener#authorization_endpoint LbListener#authorization_endpoint}.
	AuthorizationEndpoint *string `field:"required" json:"authorizationEndpoint" yaml:"authorizationEndpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/lb_listener#client_id LbListener#client_id}.
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/lb_listener#client_secret LbListener#client_secret}.
	ClientSecret *string `field:"required" json:"clientSecret" yaml:"clientSecret"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/lb_listener#issuer LbListener#issuer}.
	Issuer *string `field:"required" json:"issuer" yaml:"issuer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/lb_listener#token_endpoint LbListener#token_endpoint}.
	TokenEndpoint *string `field:"required" json:"tokenEndpoint" yaml:"tokenEndpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/lb_listener#user_info_endpoint LbListener#user_info_endpoint}.
	UserInfoEndpoint *string `field:"required" json:"userInfoEndpoint" yaml:"userInfoEndpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/lb_listener#authentication_request_extra_params LbListener#authentication_request_extra_params}.
	AuthenticationRequestExtraParams *map[string]*string `field:"optional" json:"authenticationRequestExtraParams" yaml:"authenticationRequestExtraParams"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/lb_listener#on_unauthenticated_request LbListener#on_unauthenticated_request}.
	OnUnauthenticatedRequest *string `field:"optional" json:"onUnauthenticatedRequest" yaml:"onUnauthenticatedRequest"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/lb_listener#scope LbListener#scope}.
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/lb_listener#session_cookie_name LbListener#session_cookie_name}.
	SessionCookieName *string `field:"optional" json:"sessionCookieName" yaml:"sessionCookieName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/lb_listener#session_timeout LbListener#session_timeout}.
	SessionTimeout *float64 `field:"optional" json:"sessionTimeout" yaml:"sessionTimeout"`
}

