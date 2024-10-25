// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakerworkforce


type SagemakerWorkforceOidcConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.73.0/docs/resources/sagemaker_workforce#authorization_endpoint SagemakerWorkforce#authorization_endpoint}.
	AuthorizationEndpoint *string `field:"required" json:"authorizationEndpoint" yaml:"authorizationEndpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.73.0/docs/resources/sagemaker_workforce#client_id SagemakerWorkforce#client_id}.
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.73.0/docs/resources/sagemaker_workforce#client_secret SagemakerWorkforce#client_secret}.
	ClientSecret *string `field:"required" json:"clientSecret" yaml:"clientSecret"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.73.0/docs/resources/sagemaker_workforce#issuer SagemakerWorkforce#issuer}.
	Issuer *string `field:"required" json:"issuer" yaml:"issuer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.73.0/docs/resources/sagemaker_workforce#jwks_uri SagemakerWorkforce#jwks_uri}.
	JwksUri *string `field:"required" json:"jwksUri" yaml:"jwksUri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.73.0/docs/resources/sagemaker_workforce#logout_endpoint SagemakerWorkforce#logout_endpoint}.
	LogoutEndpoint *string `field:"required" json:"logoutEndpoint" yaml:"logoutEndpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.73.0/docs/resources/sagemaker_workforce#token_endpoint SagemakerWorkforce#token_endpoint}.
	TokenEndpoint *string `field:"required" json:"tokenEndpoint" yaml:"tokenEndpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.73.0/docs/resources/sagemaker_workforce#user_info_endpoint SagemakerWorkforce#user_info_endpoint}.
	UserInfoEndpoint *string `field:"required" json:"userInfoEndpoint" yaml:"userInfoEndpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.73.0/docs/resources/sagemaker_workforce#authentication_request_extra_params SagemakerWorkforce#authentication_request_extra_params}.
	AuthenticationRequestExtraParams *map[string]*string `field:"optional" json:"authenticationRequestExtraParams" yaml:"authenticationRequestExtraParams"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.73.0/docs/resources/sagemaker_workforce#scope SagemakerWorkforce#scope}.
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
}

