// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cognitoresourceserver


type CognitoResourceServerScope struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/cognito_resource_server#scope_description CognitoResourceServer#scope_description}.
	ScopeDescription *string `field:"required" json:"scopeDescription" yaml:"scopeDescription"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/cognito_resource_server#scope_name CognitoResourceServer#scope_name}.
	ScopeName *string `field:"required" json:"scopeName" yaml:"scopeName"`
}

