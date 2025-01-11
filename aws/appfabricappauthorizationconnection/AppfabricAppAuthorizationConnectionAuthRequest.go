// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appfabricappauthorizationconnection


type AppfabricAppAuthorizationConnectionAuthRequest struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.1/docs/resources/appfabric_app_authorization_connection#code AppfabricAppAuthorizationConnection#code}.
	Code *string `field:"required" json:"code" yaml:"code"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.1/docs/resources/appfabric_app_authorization_connection#redirect_uri AppfabricAppAuthorizationConnection#redirect_uri}.
	RedirectUri *string `field:"required" json:"redirectUri" yaml:"redirectUri"`
}

