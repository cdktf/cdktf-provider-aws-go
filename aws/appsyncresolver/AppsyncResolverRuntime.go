// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appsyncresolver


type AppsyncResolverRuntime struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/resources/appsync_resolver#name AppsyncResolver#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/resources/appsync_resolver#runtime_version AppsyncResolver#runtime_version}.
	RuntimeVersion *string `field:"required" json:"runtimeVersion" yaml:"runtimeVersion"`
}

